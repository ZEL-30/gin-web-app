package assembler

var (
	UserAsm UserAssembler
)

func init() {
	UserAsm = NewUserAssembler()
}
