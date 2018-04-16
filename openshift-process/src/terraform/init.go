package terraform


type TerraformState struct {
	inited bool
	Dir string
}

var State TerraformState


func InitTerraform(dir string) {
}


