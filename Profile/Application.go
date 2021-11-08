package Profile

import (
	errorsConst "2021_2_GORYACHIE_MEKSIKANSI/Errors"
	"2021_2_GORYACHIE_MEKSIKANSI/Interfaces"
	utils "2021_2_GORYACHIE_MEKSIKANSI/Utils"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/joho/godotenv"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
	"time"
)

type Profile struct {
	DB Interfaces.WrapperProfile
}

func (p *Profile) GetProfile(id int) (*utils.Profile, error) {
	role, err := p.DB.GetRoleById(id)
	if err != nil {
		return nil, err
	}

	var result *utils.Profile
	switch role {
	case "client":
		result, err = p.DB.GetProfileClient(id)
	case "courier":
		result, err = p.DB.GetProfileCourier(id)
	case "host":
		result, err = p.DB.GetProfileHost(id)
	default:
		return nil, &errorsConst.Errors{
			Text: errorsConst.PGetProfileUnknownRole,
			Time: time.Now(),
		}
	}
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (p *Profile) UpdateName(id int, newName string) error {
	err := p.DB.UpdateName(id, newName)
	if err != nil {
		return err
	}
	return nil
}

func (p *Profile) UpdateEmail(id int, newEmail string) error {
	err := p.DB.UpdateEmail(id, newEmail)
	if err != nil {
		return err
	}
	return nil
}

func (p *Profile) UpdatePassword(id int, newPassword string) error {
	err := p.DB.UpdatePassword(id, newPassword)
	if err != nil {
		return err
	}
	return nil
}

func (p *Profile) UpdatePhone(id int, newPhone string) error {
	err := p.DB.UpdatePhone(id, newPhone)
	if err != nil {
		return err
	}
	return nil
}

func (p *Profile) UpdateAvatar(id int, newAvatar *utils.UpdateAvatar) error {
	LoadEnv()
	sess := ConnectAws()
	uploader := s3manager.NewUploader(sess)
	MyBucket := GetEnvWithKey("BUCKET_NAME")
	header := newAvatar.FileHeader
	fileNameTests := strings.Split(header.Filename, ".")
	n := len(fileNameTests)
	extensionFile := "." + fileNameTests[n - 1]
	fileName := strconv.Itoa(utils.RandomInteger(0, math.MaxInt64))
	fileResult := fileName + extensionFile

	file, errTet := header.Open()
	if errTet != nil {
		println("Не открылся file header")
		return nil
	}

	up, err := uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(MyBucket),
		ACL:    aws.String("public-read"),
		Key:    aws.String(fileResult),
		Body:   file,
	})
	if err != nil {
		println("Uploader err")
		println(up)
	}

	newAvatar.Avatar = "https://img.hmeats.fra1.digitaloceanspaces.com/" + fileResult

	err = p.DB.UpdateAvatar(id, newAvatar.Avatar)
	if err != nil {
		return err
	}
	return nil
}

func (p *Profile) UpdateBirthday(id int, newBirthday time.Time) error {
	err := p.DB.UpdateBirthday(id, newBirthday)
	if err != nil {
		return err
	}
	return nil
}

func (p *Profile) UpdateAddress(id int, newAddress utils.AddressCoordinates) error {
	err := p.DB.UpdateAddress(id, newAddress)
	if err != nil {
		return err
	}
	return nil
}

func LoadEnv() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
		os.Exit(1)
	}
}

func GetEnvWithKey(key string) string {
	return os.Getenv(key)
}

func ConnectAws() *session.Session {
	AccessKeyID := GetEnvWithKey("AWS_ACCESS_KEY_ID")
	SecretAccessKey := GetEnvWithKey("AWS_SECRET_ACCESS_KEY")
	MyRegion := GetEnvWithKey("AWS_REGION")
	sess, err := session.NewSession(
		&aws.Config{
			Endpoint: aws.String("fra1.digitaloceanspaces.com"),
			Region: aws.String(MyRegion),
			Credentials: credentials.NewStaticCredentials(
				AccessKeyID,
				SecretAccessKey,
				"",
			),
		})
	if err != nil {
		panic(err)
	}
	return sess
}