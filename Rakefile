desc "Print reminder about eating more fruit."
task :install do
	`go get -u github.com/go-swagger/go-swagger/cmd/swagger`
end

task :swagger do
	`set GO111MODULE=off `
	`swagger generate spec -o ./api/swagger.json --scan-models`
end