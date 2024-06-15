package httprsp

import "net/http"

type Writer interface {
	Write(writer http.ResponseWriter) error
}

type WriterFunc func(writer http.ResponseWriter) error

func (f WriterFunc) Write(writer http.ResponseWriter) error {
	return f(writer)
}

type DataWriter interface {
	BodyWriter
	HeaderSetter
}

type BodyWriter interface {
	WriteBody(writer http.ResponseWriter) error
}

type HeaderSetter interface {
	SetHeader(writer http.ResponseWriter) error
}

type DataWriterImpl struct {
	BodyWriter
	HeaderSetter
}

func NewDataWriter(bodyWriter BodyWriter, headerSetter HeaderSetter) DataWriterImpl {
	return DataWriterImpl{
		BodyWriter:   bodyWriter,
		HeaderSetter: headerSetter,
	}
}

type BodyWriterFunc func(writer http.ResponseWriter) error

func (f BodyWriterFunc) WriteBody(writer http.ResponseWriter) error {
	return f(writer)
}

type HeaderSetterFunc func(writer http.ResponseWriter) error

func (f HeaderSetterFunc) SetHeader(writer http.ResponseWriter) error {
	return f(writer)
}
