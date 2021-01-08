Feature: Validate API responses
  FORMACION_ACADEMICA_CRUD
  probe JSON reponses


Scenario Outline: To probe route code response  /formacion_academica
  When I send "<method>" request to "<route>" where body is json "<bodyreq>"
  Then the response code should be "<codres>"      

  Examples: 
  |method|route                  |bodyreq               |codres       |
  |GET   |/v1/formacion_academica|./files/req/Vacio.json|200 OK       |
  |GET   |/v1/formacion_academic |./files/req/Vacio.json|404 Not Found|
  |POST  |/v1/formacion_academic |./files/req/Vacio.json|404 Not Found|
  |PUT   |/v1/formacion_academic |./files/req/Vacio.json|404 Not Found|
  |DELETE|/v1/formacion_academic |./files/req/Vacio.json|404 Not Found|


Scenario Outline: To probe response route /formacion_academica       
  When I send "<method>" request to "<route>" where body is json "<bodyreq>"
  Then the response code should be "<codres>"      
  And the response should match json "<bodyres>"

  Examples: 
  |method|route                  |bodyreq               |codres         |bodyres                 |
  |GET   |/v1/formacion_academica|./files/req/Vacio.json|200 OK         |./files/res2/Vok1.json  |
  |POST  |/v1/formacion_academica|./files/req/Y1t1.json |201 Created    |./files/res2/Vok2.json  |
  |POST  |/v1/formacion_academica|./files/req/Vacio.json|400 Bad Request|./files/res2/Ierr1.json |
  |POST  |/v1/formacion_academica|./files/req/N1t1.json |400 Bad Request|./files/res2/Ierr2.json |
  |POST  |/v1/formacion_academica|./files/req/N1t2.json |400 Bad Request|./files/res2/Ierr3.json |
  |POST  |/v1/formacion_academica|./files/req/N1t3.json |400 Bad Request|./files/res2/Ierr4.json |
  |POST  |/v1/formacion_academica|./files/req/N1t4.json |400 Bad Request|./files/res2/Ierr5.json |
  |POST  |/v1/formacion_academica|./files/req/N1t5.json |400 Bad Request|./files/res2/Ierr6.json |
  |POST  |/v1/formacion_academica|./files/req/N1t6.json |400 Bad Request|./files/res2/Ierr7.json |
  |POST  |/v1/formacion_academica|./files/req/N1t7.json |400 Bad Request|./files/res2/Ierr8.json |
  |POST  |/v1/formacion_academica|./files/req/Y1t2.json |400 Bad Request|./files/res2/Ierr9.json | 
  |PUT   |/v1/formacion_academica|./files/req/Y1t2.json |200 OK         |./files/res2/Vok2.json  |
  |GETID |/v1/formacion_academica|./files/req/Vacio.json|200 OK         |./files/res2/Vok2.json  |
  |DELETE|/v1/formacion_academica|./files/req/Vacio.json|200 OK         |./files/res2/Ino.json   |
  |DELETE|/v1/formacion_academica|./files/req/Vacio.json|404 Not Found  |./files/res2/Ierr10.json|
