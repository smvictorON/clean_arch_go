package presenters

import (
	"reflect"
)

type WipeIdPresenter struct{}

func NewWipeIdPresenter() *WipeIdPresenter {
	return &WipeIdPresenter{}
}

func (pres *WipeIdPresenter) Format(data interface{}) interface{} {
	val := reflect.ValueOf(data)

	// Verifica se 'data' é um slice
	if val.Kind() != reflect.Slice {
		return data
	}

	// Cria um novo slice para armazenar os elementos modificados
	elemType := val.Type().Elem()
	newSlice := reflect.MakeSlice(reflect.SliceOf(removeFieldById(elemType, "Id")), 0, val.Len())

	// Itera sobre cada elemento do slice
	for i := 0; i < val.Len(); i++ {
		elem := val.Index(i)

		// Verifica se o elemento é um ponteiro e desreferencia
		if elem.Kind() == reflect.Ptr {
			elem = elem.Elem()
		}

		// Verifica se o elemento é uma struct
		if elem.Kind() == reflect.Struct {
			// Cria uma nova instância da struct com o novo tipo
			newElem := reflect.New(removeFieldById(elem.Type(), "Id")).Elem()

			// Copia todos os campos exceto 'Id' para a nova struct
			copyFields(elem, newElem)

			// Adiciona a nova struct ao slice
			newSlice = reflect.Append(newSlice, newElem)
		} else {
			// Se o elemento não for uma struct, mantém o elemento original
			newSlice = reflect.Append(newSlice, elem)
		}
	}

	return newSlice.Interface()
}

// removeFieldById cria um novo tipo sem o campo com nome 'Id'
func removeFieldById(structType reflect.Type, fieldName string) reflect.Type {
	var fields []reflect.StructField

	// Itera sobre os campos da struct original
	for i := 0; i < structType.NumField(); i++ {
		field := structType.Field(i)
		if field.Name != fieldName {
			fields = append(fields, field)
		}
	}

	// Cria um novo tipo struct com os campos filtrados
	return reflect.StructOf(fields)
}

// copyFields copia os valores de campos de uma struct para outra
func copyFields(src, dst reflect.Value) {
	for i := 0; i < src.NumField(); i++ {
		field := src.Type().Field(i)
		if dst.FieldByName(field.Name).IsValid() {
			dst.FieldByName(field.Name).Set(src.Field(i))
		}
	}
}
