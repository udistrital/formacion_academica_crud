package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"
	"reflect"
	"strconv"
	"strings"
	"sync"
	"testing"
	"time"

	"github.com/DATA-DOG/godog"
	"github.com/DATA-DOG/godog/colors"
	"github.com/astaxie/beego"
	"github.com/xeipuuv/gojsonschema"
)

// @resStatus codigo de respuesta a las solicitudes a la api
var resStatus string

// @resDuplex JSON repuesta Duplicacion
var resDuplex string

// @resDelete JSON repuesta Delete
var resDelete string

//@resBody JSON de respuesta a las solicitudesde la api
var resBody []byte

var savepostres map[string]interface{}

var Id float64

//@FormacionAcademica estructura de la tabla formacion_academica
type FormacionAcademica struct {
	Persona           int
	FechaInicio       time.Time
	FechaFinalizacion time.Time
	Titulacion        int
	FechaCreacion     time.Time
	FechaModificacion time.Time
}

//@FormacionAcademicaPost estructura de la tabla formacion_academica para Post de dependientes
type FormacionAcademicaPost struct {
	Id int
}

//@DatoAdicionalFormacionAcademica estructura de la tabla formacion_academica
type DatoAdicionalFormacionAcademica struct {
	FormacionAcademica FormacionAcademicaPost
	TipoDatoAdicional  int
	Valor              string
	Activo             bool
	FechaCreacion      time.Time
	FechaModificacion  time.Time
}

//@SoporteFormacionAcademica estructura de la tabla formacion_academica
type SoporteFormacionAcademica struct {
	FormacionAcademica FormacionAcademicaPost
	Documento          int
	Descripcion        string
	FechaCreacion      time.Time
	FechaModificacion  time.Time
}

//@opt opciones de godog
var opt = godog.Options{Output: colors.Colored(os.Stdout)}

//@especificacion estructura de la fecha
const especificacion = "Jan 2, 2006 at 3:04pm (MST)"

//@TestMain para realizar la ejecucion con el comando go test ./test
func TestMain(m *testing.M) {
	status := godog.RunWithOptions("godogs", func(s *godog.Suite) {
		FeatureContext(s)
	}, godog.Options{
		Format: "progress",
		Paths:  []string{"features"},
		//Randomize: time.Now().UTC().UnixNano(), // randomize scenario execution order
	})

	if st := m.Run(); st > status {
		status = st
	}
	os.Exit(status)
}

//@init inicia la aplicacion para realizar los test
func init() {
	gen_files()
	run_bee()
	//pasa las banderas al comando godog
	godog.BindFlags("godog.", flag.CommandLine, &opt)
}

//@gen_files genera los archivos de ejemplos
func gen_files() {
	t := time.Now()

	atributo1 := FormacionAcademica{
		Persona:           0,
		FechaInicio:       t,
		FechaFinalizacion: t,
		Titulacion:        0,
		FechaCreacion:     t,
		FechaModificacion: t,
	}
	rankingsJson1, _ := json.Marshal(atributo1)
	ioutil.WriteFile("./files/req/Y1t1.json", rankingsJson1, 0644)

	atributo2 := DatoAdicionalFormacionAcademica{
		FormacionAcademica: FormacionAcademicaPost{
			Id: 1,
		},
		TipoDatoAdicional: 0,
		Valor:             "string",
		Activo:            true,
		FechaCreacion:     t,
		FechaModificacion: t,
	}
	rankingsJson2, _ := json.Marshal(atributo2)
	ioutil.WriteFile("./files/req/Y2t1.json", rankingsJson2, 0644)

	atributo3 := SoporteFormacionAcademica{
		FormacionAcademica: FormacionAcademicaPost{
			Id: 1,
		},
		Documento:         0,
		Descripcion:       "string",
		FechaCreacion:     t,
		FechaModificacion: t,
	}
	rankingsJson3, _ := json.Marshal(atributo3)
	ioutil.WriteFile("./files/req/Y3t1.json", rankingsJson3, 0644)

}

//@run_bee activa el servicio de la api para realizar los test
func run_bee() {
	parametros := "FORMACION_ACADEMICA_HTTP_PORT=" + beego.AppConfig.String("httpport") + " FORMACION_ACADEMICA_CRUD__PGUSER=" + beego.AppConfig.String("PGuser") + " FORMACION_ACADEMICA_CRUD__PGPASS=" + beego.AppConfig.String("PGpass") + " FORMACION_ACADEMICA_CRUD__PGURLS=" + beego.AppConfig.String("PGurls") + " FORMACION_ACADEMICA_CRUD__PGDB=" + beego.AppConfig.String("PGdb") + " FORMACION_ACADEMICA_CRUD__SCHEMA=" + beego.AppConfig.String("PGschemas") + " bee run"
	file, err := os.Create("script.sh")
	if err != nil {
		log.Fatal("Cannot create file", err)
	}
	defer file.Close()
	fmt.Fprintln(file, "cd ..")
	fmt.Fprintln(file, parametros)

	wg := new(sync.WaitGroup)
	commands := []string{"sh script.sh &"}
	for _, str := range commands {
		wg.Add(1)
		go exe_cmd(str, wg)
	}
	time.Sleep(5 * time.Second)
	deleteFile("script.sh")
	wg.Done()
}

func deleteFile(path string) {
	// delete file
	err := os.Remove(path)
	if err != nil {
		fmt.Errorf("no se pudo eliminar el archivo")
	}
}

//@exe_cmd ejecuta comandos en la terminal
func exe_cmd(cmd string, wg *sync.WaitGroup) {
	//fmt.Println(cmd)
	parts := strings.Fields(cmd)
	out, err := exec.Command(parts[0], parts[1]).Output()
	if err != nil {
		fmt.Println("error occured")
		fmt.Printf("%s", err)
	}
	fmt.Printf("%s", out)
	wg.Done()
}

//@AreEqualJSON comparar dos JSON si son iguales retorna true de lo contrario false
func AreEqualJSON(s1, s2 string) (bool, error) {
	var o1 interface{}
	var o2 interface{}

	var err error
	err = json.Unmarshal([]byte(s1), &o1)
	if err != nil {
		return false, fmt.Errorf("Error mashalling string 1 :: %s", err.Error())
	}
	err = json.Unmarshal([]byte(s2), &o2)
	if err != nil {
		return false, fmt.Errorf("Error mashalling string 2 :: %s", err.Error())
	}

	return reflect.DeepEqual(o1, o2), nil
}

//@toJson convierte string en JSON
func toJson(p interface{}) string {
	bytes, err := json.Marshal(p)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	return string(bytes)
}

//@getPages convierte en un tipo el json
func getPages(ruta string) []byte {
	raw, err := ioutil.ReadFile(ruta)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	var c []byte
	c = raw
	return c
}

//@iSendRequestToWhereBodyIsJson realiza la solicitud a la API
func iSendRequestToWhereBodyIsJson(method, endpoint, bodyreq string) error {
	var url string

	if method == "GET" || method == "POST" {
		url = "http://localhost:" + beego.AppConfig.String("httpport") + endpoint
	} else {
		if method == "PUT" || method == "DELETE" {
			str := strconv.FormatFloat(Id, 'f', 0, 64)
			url = "http://localhost:" + beego.AppConfig.String("httpport") + endpoint + "/" + str
		}
	}
	if method == "GETID" {
		method = "GET"
		str := strconv.FormatFloat(Id, 'f', 0, 64)
		url = "http://localhost:" + beego.AppConfig.String("httpport") + endpoint + "/" + str
	}
	if method == "DELETE" {
		str := strconv.FormatFloat(Id, 'f', 0, 64)
		url = "http://localhost:" + beego.AppConfig.String("httpport") + endpoint + "/" + str
		resDelete = "{\"Id\":" + str + "}"
		ioutil.WriteFile("./files/res1/Ino.json", []byte(resDelete), 0644)
		ioutil.WriteFile("./files/res2/Ino.json", []byte(resDelete), 0644)
		ioutil.WriteFile("./files/res3/Ino.json", []byte(resDelete), 0644)
	}

	pages := getPages(bodyreq)

	req, err := http.NewRequest(method, url, bytes.NewBuffer(pages))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	bodyr, _ := ioutil.ReadAll(resp.Body)

	resStatus = resp.Status
	resBody = bodyr

	if method == "POST" && resStatus == "201 Created" {
		if endpoint == "/v1/dato_adicional_formacion_academica" {
			ioutil.WriteFile("./files/req/Y2t2.json", resBody, 0644)
		} else if endpoint == "/v1/soporte_formacion_academica" {
			ioutil.WriteFile("./files/req/Y3t2.json", resBody, 0644)
		} else {
			ioutil.WriteFile("./files/req/Y1t2.json", resBody, 0644)
		}
		json.Unmarshal([]byte(bodyr), &savepostres)
		Id = savepostres["Id"].(float64)

		resDuplex = "{\r\n\t\"Development\": null,\r\n\t\"Message\": \"The request contains incorrect syntax\",\r\n\t\"Status\": \"400\",\r\n\t\"System\": {\r\n"
		resDuplex = resDuplex + "\t\t\"Severity\": \"ERROR\",\r\n\t\t\"Code\": \"23505\",\r\n\t\t\"Message\": \"llave duplicada viola restricción de unicidad «pk_dato_adicional_formacion_academica»\",\r\n"
		resDuplex = resDuplex + "\t\t\"Detail\": \"Ya existe la llave (id)=(" + strconv.Itoa(int(Id)) + ").\",\r\n\t\t\"Hint\": \"\",\r\n\t\t\"Position\": \"\",\r\n\t\t\"InternalPosition\": \"\",\r\n"
		resDuplex = resDuplex + "\t\t\"InternalQuery\": \"\",\r\n\t\t\"Where\": \"\",\r\n\t\t\"Schema\": \"formacion_academica\",\r\n\t\t\"Table\": \"dato_adicional_formacion_academica\",\r\n\t\t\"Column\": \"\",\r\n"
		resDuplex = resDuplex + "\t\t\"DataTypeName\": \"\",\r\n\t\t\"Constraint\": \"pk_dato_adicional_formacion_academica\",\r\n\t\t\"File\": \"nbtinsert.c\",\r\n\t\t\"Line\": \"434\",\r\n\t\t\"Routine\": \"_bt_check_unique\"\r\n\t}\r\n}"

		ioutil.WriteFile("./files/res1/Ierr9.json", []byte(resDuplex), 0644)

		resDuplex = "{\r\n\t\"Development\": null,\r\n\t\"Message\": \"The request contains incorrect syntax\",\r\n\t\"Status\": \"400\",\r\n\t\"System\": {\r\n"
		resDuplex = resDuplex + "\t\t\"Severity\": \"ERROR\",\r\n\t\t\"Code\": \"23505\",\r\n\t\t\"Message\": \"llave duplicada viola restricción de unicidad «pk_formacion_academica»\",\r\n"
		resDuplex = resDuplex + "\t\t\"Detail\": \"Ya existe la llave (id)=(" + strconv.Itoa(int(Id)) + ").\",\r\n\t\t\"Hint\": \"\",\r\n\t\t\"Position\": \"\",\r\n\t\t\"InternalPosition\": \"\",\r\n"
		resDuplex = resDuplex + "\t\t\"InternalQuery\": \"\",\r\n\t\t\"Where\": \"\",\r\n\t\t\"Schema\": \"formacion_academica\",\r\n\t\t\"Table\": \"formacion_academica\",\r\n\t\t\"Column\": \"\",\r\n"
		resDuplex = resDuplex + "\t\t\"DataTypeName\": \"\",\r\n\t\t\"Constraint\": \"pk_formacion_academica\",\r\n\t\t\"File\": \"nbtinsert.c\",\r\n\t\t\"Line\": \"434\",\r\n\t\t\"Routine\": \"_bt_check_unique\"\r\n\t}\r\n}"

		ioutil.WriteFile("./files/res2/Ierr9.json", []byte(resDuplex), 0644)

		resDuplex = "{\r\n\t\"Development\": null,\r\n\t\"Message\": \"The request contains incorrect syntax\",\r\n\t\"Status\": \"400\",\r\n\t\"System\": {\r\n"
		resDuplex = resDuplex + "\t\t\"Severity\": \"ERROR\",\r\n\t\t\"Code\": \"23505\",\r\n\t\t\"Message\": \"llave duplicada viola restricción de unicidad «pk_soporte_formacion_academica»\",\r\n"
		resDuplex = resDuplex + "\t\t\"Detail\": \"Ya existe la llave (id)=(" + strconv.Itoa(int(Id)) + ").\",\r\n\t\t\"Hint\": \"\",\r\n\t\t\"Position\": \"\",\r\n\t\t\"InternalPosition\": \"\",\r\n"
		resDuplex = resDuplex + "\t\t\"InternalQuery\": \"\",\r\n\t\t\"Where\": \"\",\r\n\t\t\"Schema\": \"formacion_academica\",\r\n\t\t\"Table\": \"soporte_formacion_academica\",\r\n\t\t\"Column\": \"\",\r\n"
		resDuplex = resDuplex + "\t\t\"DataTypeName\": \"\",\r\n\t\t\"Constraint\": \"pk_soporte_formacion_academica\",\r\n\t\t\"File\": \"nbtinsert.c\",\r\n\t\t\"Line\": \"434\",\r\n\t\t\"Routine\": \"_bt_check_unique\"\r\n\t}\r\n}"

		ioutil.WriteFile("./files/res3/Ierr8.json", []byte(resDuplex), 0644)
	}
	return nil
}

//@theResponseCodeShouldBe valida el codigo de respuesta
func theResponseCodeShouldBe(arg1 string) error {
	if resStatus != arg1 {
		return fmt.Errorf("se esperaba el codigo de respuesta .. %s .. y se obtuvo el codigo de respuesta .. %s .. ", arg1, resStatus)
	}
	return nil
}

//@theResponseShouldMatchJson valida el JSON de respuesta
func theResponseShouldMatchJson(arg1 string) error {
	div := strings.Split(arg1, "")

	pages := getPages(arg1)
	//areEqual, _ := AreEqualJSON(string(pages), string(resBody))
	if div[13] == "V" {
		schemaLoader := gojsonschema.NewStringLoader(string(pages))
		documentLoader := gojsonschema.NewStringLoader(string(resBody))
		result, err := gojsonschema.Validate(schemaLoader, documentLoader)
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}
		if result.Valid() {
			return nil
		} else {
			return fmt.Errorf("Errores : %s", result.Errors())

			return nil
		}
	}
	if div[13] == "I" {
		areEqual, _ := AreEqualJSON(string(pages), string(resBody))
		if areEqual {
			return nil
		} else {
			return fmt.Errorf(" se esperaba el body de respuesta %s y se obtuvo %s", string(pages), resBody)
		}

	}
	return nil
}

func FeatureContext(s *godog.Suite) {
	s.Step(`^I send "([^"]*)" request to "([^"]*)" where body is json "([^"]*)"$`, iSendRequestToWhereBodyIsJson)
	s.Step(`^the response code should be "([^"]*)"$`, theResponseCodeShouldBe)
	s.Step(`^the response should match json "([^"]*)"$`, theResponseShouldMatchJson)
}
