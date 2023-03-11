package postgres

import (
	"fmt"
	"log"
	"os"

	"go-grpc-crud/server/domain"
	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)


type postgresRepository struct {
	db        *gorm.DB
	tableName string
}

type Settings struct {
	host     string
	dbname     string
	user    string
	password string
	port     string
}

func InitializeSettings() Settings {
	host := os.Getenv("host")
	dbname := os.Getenv("dbname")
	user := os.Getenv("user")
	port := os.Getenv("port")
	password := os.Getenv("password")

	switch {
	case host == "":
		fmt.Println("Environment variable host not set.")
		os.Exit(1)
	case dbname == "":
		fmt.Println("Environment variable dbname not set.")
		os.Exit(1)
	case user == "":
		fmt.Println("Environment variable user not set.")
		os.Exit(1)
	case password == "":
		fmt.Println("Environment variable password not set.")
		os.Exit(1)
	case port == "":
		fmt.Println("Environment variable port not set.")
		os.Exit(1)
	}

	settings := Settings{
		host:     host,
		port:     port,
		dbname:     dbname,
		user:     user,
		password: password,
	}

	return settings
}

func NewPostgresRepository() (domain.MovieRepository, error) {
	settings := InitializeSettings()
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	db, err  := gorm.Open(fmt.Sprintf("postgres://%v:%v@%v:%v/%v?sslmode=disable", settings.user, settings.password, settings.host, settings.port, settings.dbname))
	if err != nil {
		log.Fatal(err)
	}
	
	db.LogMode(false)
	db.AutoMigrate(&domain.Movie{})
	repo := &postgresRepository{}
	repo.db = db
	repo.tableName = "Movies"

	return repo, nil
}
func (repo *postgresRepository) CreateMovie(movie *domain.Movie) (*domain.Movie, error) {

	if err := repo.db.Create(&movie).Error; err != nil {
		return nil, err
	}
	return movie, nil
}

func (repo *postgresRepository) GetMovie(id string) (*domain.Movie, error) {
	movie := domain.Movie{}
	if result := repo.db.Find(&movie); result.Error != nil {
		return nil, result.Error
	}
	return &movie, nil
}

func (repo *postgresRepository) GetMovies() (*[]domain.Movie, error) {
	movies := []domain.Movie{}
	if result := repo.db.Find(&movies); result.Error != nil {
		fmt.Println(result.Error)
		return nil, result.Error
	}
	return &movies, nil
}

func (repo *postgresRepository) UpdateMovie(movie *domain.Movie) (*domain.Movie, error) {
	if result := repo.db.Save(movie); result.Error != nil {

		return nil, result.Error
	}
	return movie, nil

}

func (repo *postgresRepository) DeleteMovie(id string) error {
	movie := domain.Movie{}
	if result := repo.db.Find(&movie); result.Error != nil {
		return result.Error
	}
	fmt.Println(id)
	if err := repo.db.Where("id = ? ", id).Delete(&movie).Error; err != nil {

		return err
	}
	return nil

}
