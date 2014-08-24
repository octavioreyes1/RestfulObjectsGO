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
   	serv.ResponseBuilder().CacheMaxAge(60 * 60 * 24) //List cacheable for a day. More work to come on this, Etag, etc
   	return declaranteStore
}
////// VER UN DECLARANTE
func (serv DeclaranetService) ViewDeclarante(id int) (retDeclarante Declarante) {
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
   	//Item Id not in database, so create new
   	d.Id = len(declaranteStore)+1
   	declaranteStore = append(declaranteStore, d)
       serv.ResponseBuilder().Created("http://localhost:8787/declarante-service/declarantes/" + string(d.Id)) //Created, http 201
}
//////// MODIFICAR DATOS DE UN DECLARANTE
func (serv DeclaranetService) ModDeclarante(d Declarante) {
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