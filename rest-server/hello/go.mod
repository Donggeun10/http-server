module example.com/hello

go 1.22.5

replace example.com/greetings => ../greetings

require (
	example.com/crashReportFlatbuffers v0.0.0-00010101000000-000000000000
	example.com/greetings v0.0.0-00010101000000-000000000000
	github.com/gorilla/mux v1.8.1
)

require github.com/google/flatbuffers v24.3.25+incompatible // indirect

replace example.com/crashReportFlatbuffers => ../crashReportFlatbuffers
