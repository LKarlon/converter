type ServiceTransport interface {

	FileDecode(ctx *gin.) (employeeID, tripID string, files []models.File, err error)
}