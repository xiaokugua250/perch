package auth



/**

采用go-guard进行身份认证
参考:
	https://github.com/shaj13/go-guardian
*/
import (
	"github.com/shaj13/go-guardian/v2"
	"github.com/shaj13/libcache"
	_ "github.com/shaj13/libcache/fifo"

	"github.com/shaj13/go-guardian/v2/auth"
	"github.com/shaj13/go-guardian/v2/auth/strategies/basic"
	"github.com/shaj13/go-guardian/v2/auth/strategies/jwt"
	"github.com/shaj13/go-guardian/v2/auth/strategies/union"
)



func GuardLdapSetup()(error){

}



func GuardJwtSetup()(error){

}


func GuardOauth2Setup()(error){

}


func GuardBasicSetup()(error){

}

func GuardBearTokenSetup()(error){

}

func GuardStaticTokenSetup()(error){

}