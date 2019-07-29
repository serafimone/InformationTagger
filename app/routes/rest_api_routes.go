package routes

import (
	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"github.com/serafimone/InformationTagger/app/core/authentication"
	"github.com/serafimone/InformationTagger/app/services"
	"github.com/serafimone/InformationTagger/app/utils"
)

func GetRESTAPIRoutes(router *mux.Router) *mux.Router {
	router.Handle(
		utils.APIVersion+"documents",
		negroni.New(
			negroni.HandlerFunc(services.GetAllDocuments),
		)).Methods("GET")
	router.Handle(
		utils.APIVersion+"documents/{document_id}",
		negroni.New(
			negroni.HandlerFunc(services.GetDocument),
		)).Methods("GET")
	router.Handle(
		utils.APIVersion+"documents/document",
		negroni.New(
			negroni.HandlerFunc(services.CreateDocument),
		)).Methods("POST")
	router.Handle(
		utils.APIVersion+"documents/{document_id}",
		negroni.New(
			negroni.HandlerFunc(authentication.RequireTokenAuthentication),
			negroni.HandlerFunc(services.UpdateDocumentTitle),
		)).Methods("POST")
	router.Handle(
		utils.APIVersion+"documents/{document_id}",
		negroni.New(
			negroni.HandlerFunc(services.DeleteDocument),
		)).Methods("DELETE", "OPTIONS")
	router.Handle(
		utils.APIVersion+"documents/{document_id}/records",
		negroni.New(
			negroni.HandlerFunc(authentication.RequireTokenAuthentication),
			negroni.HandlerFunc(services.GetAllDocumentRecords),
		)).Methods("GET")
	router.Handle(
		utils.APIVersion+"documents/{document_id}/records/{record_id}",
		negroni.New(
			negroni.HandlerFunc(authentication.RequireTokenAuthentication),
			negroni.HandlerFunc(services.GetDocumentRecord),
		)).Methods("GET")
	router.Handle(
		utils.APIVersion+"documents/{document_id}/records/record",
		negroni.New(
			negroni.HandlerFunc(services.AddRecordToDocument),
		)).Methods("POST")
	router.Handle(
		utils.APIVersion+"documents/{document_id}/records/{record_id}",
		negroni.New(
			negroni.HandlerFunc(authentication.RequireTokenAuthentication),
			negroni.HandlerFunc(services.UpdateDocumentRecord),
		)).Methods("POST")
	router.Handle(utils.APIVersion+"documents/{document_id}/records/{record_id}",
		negroni.New(
			negroni.HandlerFunc(services.DeleteDocumentRecord),
		)).Methods("DELETE", "OPTIONS")
	router.Handle(utils.APIVersion+"form",
		negroni.New(
			negroni.HandlerFunc(services.FormDocument),
		)).Methods("POST", "OPTIONS")
	return router
}
