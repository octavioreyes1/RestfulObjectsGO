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

// Se incluyen los servicios solicitados por el cliente AROW
// como domainTypes, services, users y version

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
   	servicios   gorest.EndPoint `method:"GET" path:"/services/" output:"Json"`
   	domain   gorest.EndPoint `method:"GET" path:"/domainTypes/" output:"Jsondomain"`
   	version   gorest.EndPoint `method:"GET" path:"/version/" output:"Jsonover"`
   	usuarios   gorest.EndPoint `method:"GET" path:"/users/" output:"Jsonuser"`

}



func (serv DeclaranetService) Self()(j Json){
    rb:= serv.ResponseBuilder()
    rb.AddHeader("Access-Control-Allow-Origin", "*")
    rb.AddHeader("Access-Control-Allow-Headers", "X-HTTP-Method-Override")
    rb.AddHeader("Access-Control-Allow-Headers", "X-Xsrf-Cookie")
    rb.AddHeader("Access-Control-Expose-Headers", "X-Xsrf-Cookie")
    link1 := Link{Rel:"self", Method:"GET", Type:"application/json; profile=\"urn:org.restfulobjects: repr-types/homepage\"; charset=utf-8"}
    link2 := Link{Rel:"user", Method:"GET", Type:"application/json;profile=\"urn:org.restfulobjects/homepage\"", Href:"http://localhost:8787/declaranet-service/user"}
    link3 := Link{Rel:"services", Method:"GET", Type:"application/json;profile=\"urn:org.restfulobjects/list\"", Href:"http://localhost:8787/declaranet-service/services"}
    link4 := Link{Rel:"version", Method:"GET", Type:"application/json;profile=\"urn:org.restfulobjects/version\"", Href:"http://localhost:8787/declaranet-service/version"}

    link10 := Link{Rel:"urn:orgp.restfulobjects:rels/declarante", Method:"GET", Type:"application/json; profile=\"urn:org.restfulobjects: repr-types/declarante\"; charset=utf-8", Href:"http://localhost:8787/declaranet-service/declarantes"}
    linkStore := []Link{link1, link2, link3, link4, link10}
    n := Json{Links:linkStore, Extensions:Ext{}}
    y  := 2
        if (y==2){
        j=n
        return
    }        
    j = n
    serv.ResponseBuilder().SetResponseCode(404).Overide(true)
    return  
}


func (serv DeclaranetService) Domain()(j Jsondomain){
    rb:= serv.ResponseBuilder()
    rb.AddHeader("Access-Control-Allow-Origin", "*")
    rb.AddHeader("Access-Control-Allow-Headers", "X-HTTP-Method-Override")
    rb.AddHeader("Access-Control-Allow-Headers", "X-Xsrf-Cookie")
    rb.AddHeader("Access-Control-Expose-Headers", "X-Xsrf-Cookie")
    link1 := Link{Rel:"self",Method:"GET", Type:"application/json;profile=\"urn:org.restfulobjects/typelist\"", Href:"http://localhost:8787/declaranet-service/domainTypes"}
    linkStore := []Link{link1}

    n := Jsondomain{Links:linkStore, Values:linkStore, Extensions:Ext{}}
    y  := 2
    if (y==2){
        j=n
        return
    }        
    j = n
    serv.ResponseBuilder().SetResponseCode(404).Overide(true)
     return  
}



func (serv DeclaranetService) Servicios()(j Json){
     rb:= serv.ResponseBuilder()
        rb.AddHeader("Access-Control-Allow-Origin", "*")
        rb.AddHeader("Access-Control-Allow-Headers", "X-HTTP-Method-Override")
        rb.AddHeader("Access-Control-Allow-Headers", "X-Xsrf-Cookie")
        rb.AddHeader("Access-Control-Expose-Headers", "X-Xsrf-Cookie")
        link1 := Link{Rel:"self",Method:"GET", Type:"application/json; profile=\"urn:org.restfulobjects:repr-types/list\"; charset=utf-8; x-ro-element-type=\"System.Object\"", Href:"http://localhost:8787/declaranet-service/services"}
        link2 := Link{Rel:"up",Method:"GET", Type:"application/json; profile=\"urn:org.restfulobjects:repr-types/homepage\"; charset=utf-8", Href:"http://localhost:8787/declaranet-service/"}
        linkStore := []Link{link1, link2}
        n := Json{Links:linkStore, Extensions:Ext{}}
        y  := 2
    
        if (y==2){
            j=n
            return
        }        
        j = n
        serv.ResponseBuilder().SetResponseCode(404).Overide(true)
     return  
}



func (serv DeclaranetService) Version()(j Jsonver){
    rb:= serv.ResponseBuilder()
    rb.AddHeader("Access-Control-Allow-Origin", "*")
    rb.AddHeader("Access-Control-Allow-Headers", "X-HTTP-Method-Override")
    rb.AddHeader("Access-Control-Allow-Headers", "X-Xsrf-Cookie")
    rb.AddHeader("Access-Control-Expose-Headers", "X-Xsrf-Cookie")
    link1 := Link{Rel:"self",Method:"GET", Type:"application/json;profile=\"urn:org.restfulobjects/version\"", Href:"http://localhost:8787/declaranet-service/version"}
    linkStore := []Link{link1}
    n := Jsonver{Links:linkStore, Specversion:"1.0", ImplVersion:"1.1.0", Extensions:Ext{}}
    y  := 2
    if (y==2){
        j=n
        return
    }        
    j = n
    serv.ResponseBuilder().SetResponseCode(404).Overide(true)
     return  
}



func (serv DeclaranetService) Usuarios()(j Jsonuser){
    rb:= serv.ResponseBuilder()
    rb.AddHeader("Access-Control-Allow-Origin", "*")
    rb.AddHeader("Access-Control-Allow-Headers", "X-HTTP-Method-Override")
    rb.AddHeader("Access-Control-Allow-Headers", "X-Xsrf-Cookie")
    rb.AddHeader("Access-Control-Expose-Headers", "X-Xsrf-Cookie")
    link1 := Link{Rel:"self",Method:"GET", Type:"application/json; profile=\"urn:org.restfulobjects:repr-types/user\"; charset=utf-8", Href:"http://localhost:8787/declaranet-service/user"}
    link2 := Link{Rel:"up",Method:"GET", Type:"application/json; profile=\"urn:org.restfulobjects:repr-types/homepage\"; charset=utf-8", Href:"http://localhost:8787/declaranet-service/"}
    linkStore := []Link{link1, link2}
    rol:=[]Rol{}
    n := Jsonuser{Links:linkStore, Extensions:Ext{}, UserName:"", Roles:rol}
    y  := 2
    if (y==2){
        j=n
        return
    }        
    j = n
    serv.ResponseBuilder().SetResponseCode(404).Overide(true)
     return  
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

/// Estructura Link
type Link struct {
   	Rel    	string
   	Method string
        Type  string
   	Href  string
}

/// Estructura Jsonuser
type Jsonuser struct {
   	Links Link[]
   	Extensions Ext
        Roles  Rol[]
   	Username  string
}

/// Estructura Jsondomain
type Jsondomain struct {
   	Links  Link[]
   	Values Link[]
        Extensions Ext
}

/// Estructura Jsonover
type Jsonover struct {
   	Links Link[]
   	Extensions  Ext
        Specversion  string
   	Implversion  string
}




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