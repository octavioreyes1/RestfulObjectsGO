package main
// Librerías necesarias
import (
	"code.google.com/p/gorest"
   	"net/http"
    	"strconv"
)

// Programa para generar los servicios GET, PUT, POST y DELETE
// (Altas, bajas, cambios y consultas)
// de un conjunto de estructuras con el nombre de Declarante

// Se agregaron los headers para que cualquier cliente en la misma red
// pueda consumir los servicios de este programa

// Función principal
func main(){
   gorest.RegisterService(new(DeclaranetService))
   http.Handle("/",gorest.Handle())	
  http.ListenAndServe(":8787",nil)
}
/// Estructura principal que contiene la lista de servicios GET, POST, PUT y DELETE
type DeclaranetService struct {
   	//Service level config
   	gorest.RestService `root:"/declaranet-service/" consumes:"application/json" produces:"application/json"`

   	//End-Point level configs: Field names must be the same as the corresponding method names,
   	// but not-exported (starts with lowercase)

  	
           viewDeclarante   gorest.EndPoint `method:"GET" path:"/declarantes/{Id:int}" output:"Declarante"`
   	listDeclarantes   gorest.EndPoint `method:"GET" path:"/declarantes/" output:"[]Declarante"`
   	deleteDeclarante gorest.EndPoint `method:"DELETE" path:"/declarantes/{Id:int}"`
   	addDeclarante 	gorest.EndPoint `method:"PUT" path:"/declarantes/" postdata:"Declarante"`
   	modDeclarante 	gorest.EndPoint `method:"POST" path:"/declarantes/" postdata:"Declarante"`
}
////LISTAR DECLARANTES
func (serv DeclaranetService) ListDeclarantes() []Declarante {
        rb:=serv.ResponseBuilder()
        rb.AddHeader("Access-Control-Allow-Origin", "*")
        rb.AddHeader("Access-Control-Allow-Headers", "X-HTTP-Method-Override")
        rb.AddHeader("Access-Control-Allow-Headers", "X-Xsrf-Cookie")
        rb.AddHeader("Access-Control-Expose-Headers", "X-Xsrf-Cookie").

   	serv.ResponseBuilder().CacheMaxAge(60 * 60 * 24) //List cacheable for a day. More work to come on this, Etag, etc
   	return declaranteStore
}
////// VER UN DECLARANTE
func (serv DeclaranetService) ViewDeclarante(id int) (retDeclarante Declarante) {
        rb:=serv.ResponseBuilder()
        rb.AddHeader("Access-Control-Allow-Origin", "*")
        rb.AddHeader("Access-Control-Allow-Headers", "X-HTTP-Method-Override")
        rb.AddHeader("Access-Control-Allow-Headers", "X-Xsrf-Cookie")
        rb.AddHeader("Access-Control-Expose-Headers", "X-Xsrf-Cookie").

   	for _, declarante := range declaranteStore {
           	if declarante.Id == id {
                   	retDeclarante = declarante
                   	return
           	}
   	}
       serv.ResponseBuilder().SetResponseCode(404).Overide(true)
   	return
}
 
/////// BORRAR UN DECLARANTE
func (serv DeclaranetService) DeleteDeclarante(id int) {
        rb:=serv.ResponseBuilder()
        rb.AddHeader("Access-Control-Allow-Origin", "*")
        rb.AddHeader("Access-Control-Allow-Headers", "X-HTTP-Method-Override")
        rb.AddHeader("Access-Control-Allow-Headers", "X-Xsrf-Cookie")
        rb.AddHeader("Access-Control-Expose-Headers", "X-Xsrf-Cookie").

   	for pos, declarante := range declaranteStore {
           	if declarante.Id == id {                	
                   	declaranteStore[pos] = declarante
                   	return //Default http code for DELETE is 200
           	}
   	}
   	serv.ResponseBuilder().SetResponseCode(404).Overide(true)
   	return
}
///////// AGREGAR UN DECLARANTE
func (serv DeclaranetService) AddDeclarante(d Declarante) {
        rb:=serv.ResponseBuilder()
        rb.AddHeader("Access-Control-Allow-Origin", "*")
        rb.AddHeader("Access-Control-Allow-Headers", "X-HTTP-Method-Override")
        rb.AddHeader("Access-Control-Allow-Headers", "X-Xsrf-Cookie")
        rb.AddHeader("Access-Control-Expose-Headers", "X-Xsrf-Cookie").

   	//Item Id not in database, so create new
   	d.Id = len(declaranteStore)+1
   	declaranteStore = append(declaranteStore, d)
       serv.ResponseBuilder().Created("http://localhost:8787/declarante-service/declarantes/" + string(d.Id)) //Created, http 201
}
//////// MODIFICAR DATOS DE UN DECLARANTE
func (serv DeclaranetService) ModDeclarante(d Declarante) {
        rb:=serv.ResponseBuilder()
        rb.AddHeader("Access-Control-Allow-Origin", "*")
        rb.AddHeader("Access-Control-Allow-Headers", "X-HTTP-Method-Override")
        rb.AddHeader("Access-Control-Allow-Headers", "X-Xsrf-Cookie")
        rb.AddHeader("Access-Control-Expose-Headers", "X-Xsrf-Cookie").

   	for pos, declarante := range declaranteStore {
           	if declarante.Id == d.Id {
                   	declaranteStore[pos] = d
                       serv.ResponseBuilder().SetResponseCode(200) //Updated http 200, or you could just return without setting this. 200 is the default for POST
                   	return
           	}
   	}
}
//**************** End of service *******************
/// Estructura Declarante y sus campos
type Declarante struct {
   	Id    	int
   	Nombre string
            ApPaterno  string
   	ApMaterno  string
}
var (
          // Declación de un arreglo de Declarantes
   	declaranteStore  []Declarante
)
// Función de inicialización para crear un arreglo de estructuras de tipo Declarante
func init() {
    	declaranteStore = make([]Declarante, 0)
    	initDeclarantes()
}
/////////// INICIALIZAR UN ARREGLO CON 10 DECLARANTES
func initDeclarantes() {
   	for i := 1; i <= 10; i++ {
           	declaranteStore = append(declaranteStore, Declarante{Id: i,
                 	Nombre: "Nombre" + strconv.Itoa(i),
                 	ApPaterno:  "Apellido P." + strconv.Itoa(i),
	             ApMaterno:  "Apellido M." + strconv.Itoa(i)})
   	}
}