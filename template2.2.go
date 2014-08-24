package main
import "text/template"
import "os"
import "reflect"
import "fmt"
import "strings"

// Programa para obtener los atributos y nombre 
// de una estructura de tipo Declarante
// por medio de reflection o reflexion.

// Estructura
type Item struct {
           	Text string
}
func main() {
           	de := reflect.TypeOf(Declarante{})
           	var ClassName string = fmt.Sprintf("%s", de.Name())
           	var className string = strings.ToLower(ClassName)
           	var parts []string
           	for name, mtype := range attributes(&Declarante{}) {
                          	if name != "Id" {
                                          	parts = append(parts, fmt.Sprintf("%s %s\n", name, mtype.Name()))
                          	}
           	}
           	var campos string = strings.Join(parts, "")
           	t, _ := template.ParseFiles("service2.2.tmpl")
           	// field names don't have to be capitalized
           	params := map[string]interface{}{"serviceName": "declaranet",
                          	"ServiceName": "Declaranet"}
           	params["ServiceName"] = ClassName
           	params["serviceName"] = className
           	params["ClassName"] = ClassName
           	params["className"] = className
           	//********************* CREAR CAMPOS DE LA ESTRUCTURA *****************************//
           	params["campos"] = campos
           	//********************* INICIALIZAR VALORES ***************************************//
           	var partsini []string
           	for i := 0; i < de.NumField(); i++ {
                          	field := de.Field(i)
                          	if field.Name != "Id" {
                                          	partsini = append(partsini, fmt.Sprintf("%s: \"%s\" + strconv.Itoa(i),\n", field.Name, field.Name))
                          	}
           	}
           	var camposini string = strings.Join(partsini, "")
           	sz := len(camposini)
           	if sz > 0 && camposini[sz-2] == ',' {
                          	camposini = camposini[:sz-2]
           	}
           	camposini = camposini + "})"
           	params["camposini"] = camposini
           	// EJEMPLO
           	//	Nombre: "Nombre" + strconv.Itoa(i),
           	//	ApPaterno:  "Apellido P." + strconv.Itoa(i),
           	//	ApMaterno:  "Apellido M." + strconv.Itoa(i)})
           	//********* REEMPLAZAR VALORES ****************
           	t.Execute(os.Stdout, params)
}
// REFLECTION
func attributes(m interface{}) map[string]reflect.Type {
           	typ := reflect.TypeOf(m)
           	if typ.Kind() == reflect.Ptr {
                          	typ = typ.Elem()
           	}
           	attrs := make(map[string]reflect.Type)
           	if typ.Kind() != reflect.Struct {
                          	fmt.Printf("%w type cant have attributes inspected\n", typ.Kind())
                          	return attrs
           	}
           	for i := 0; i < typ.NumField(); i++ {
                          	p := typ.Field(i)
                          	if !p.Anonymous {
                                          	attrs[p.Name] = p.Type
                          	}
           	}
           	return attrs
}
// ESTRUCTURA
type Declarante struct {
           	Id    	int
           	Nombre	string
           	ApPaterno string
           	ApMaterno string
           	Rfc   	string
}
var (
           	declaranteStore []Declarante
)
