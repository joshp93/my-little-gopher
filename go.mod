module my-little-gopher

go 1.20

replace my-little-gopher/hello => ./hello

replace my-little-gopher/greetings => ./greetings

require my-little-gopher/hello v0.0.0-00010101000000-000000000000

require my-little-gopher/greetings v0.0.0-00010101000000-000000000000 // indirect
