package Interface
import(
	"context"
	authProto "2021_2_GORYACHIE_MEKSIKANSI/internal/Microservices/Authorization/proto"
)

type AuthorizationService interface {
	CheckAccessUser(ctx context.Context, cookie *authProto.Defense) (*authProto.CheckAccess, error)
	NewCSRFUser(ctx context.Context, cookie *authProto.Defense) (*authProto.CSRFResponse, error)
	GetIdByCookie(ctx context.Context, cookie *authProto.Defense) (*authProto.IdClientResponse, error)
	SignUp(ctx context.Context, signup *authProto.RegistrationRequest) (*authProto.DefenseResponse, error)
	Login(ctx context.Context, login *authProto.Authorization) (*authProto.DefenseResponse, error)
	Logout(ctx context.Context, CSRF *authProto.CSRF) (*authProto.CSRFResponse, error)
}
