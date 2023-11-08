cd skill || return
zip -r ../skill.zip . -x go.mod go.sum main.go
cd ../ || return
zip skill.zip go.mod
