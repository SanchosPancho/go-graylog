package endpoint

// PipelineConnections returns a Pipeline Connections API's endpoint url.
func (ep *Endpoints) PipelineConnections() string {
	return ep.pipelineConnections
}

// PipelineConnectionsOfStream returns a Pipeline Connections for a given stream API's endpoint url.
func (ep *Endpoints) PipelineConnectionsOfStream(id string) string {
	return ep.pipelineConnections + "/" + id
}

// ConnectStreamsToPipeline returns a connect streams to a pipeline API's endpoint url.
func (ep *Endpoints) ConnectStreamsToPipeline() string {
	return ep.connectStreamsToPipeline
}

// ConnectPipelinesToStream returns a connect processing pipelines to a stream API's endpoint url.
func (ep *Endpoints) ConnectPipelinesToStream() string {
	return ep.connectPipelinesToStream
}
