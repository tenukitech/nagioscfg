package nagioscfg

//GenericObject provides a structure to map a generic nagios object type
type GenericObject struct {
	//Type represents the type of the configuration object -- host, service, etc.
	ObjectType string
	//Attributes captures all attributes of the object.
	Attributes map[string]string
}

//NewGenericObject returns a new, initialized generic object of the requested
//type.  TODO: extend this to allow validation of the object by type.
func NewGenericObject(objectType string) GenericObject {
	return GenericObject{
		ObjectType: objectType,
		Attributes: make(map[string]string),
	}
}
