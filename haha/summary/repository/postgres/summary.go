package postgres

import (
	"database/sql"
	"fmt"
	"joblessness/haha/models"
	"strings"
)

type Summary struct {
	ID uint64
	AuthorID uint64
	Keywords sql.NullString
}

type Education struct {
	SummaryID uint64
	Institution sql.NullString
	Speciality sql.NullString
	Graduated sql.NullString
	Type sql.NullString
}

type Experience struct {
	SummaryID uint64
	CompanyName sql.NullString
	Role sql.NullString
	Responsibilities sql.NullString
	Start sql.NullString
	Stop sql.NullString
}

type User struct {
	ID             uint64
	Login          sql.NullString
	Password       sql.NullString
	OrganizationID uint64
	PersonID       uint64
	Tag            sql.NullString
	Email          sql.NullString
	Phone          sql.NullString
	Registered     sql.NullTime
	Avatar         sql.NullString
}

type Person struct {
	ID       sql.NullString
	Name     sql.NullString
	Gender   sql.NullString
	Birthday sql.NullTime
}

func toPostgres(s *models.Summary) (summary *Summary, educations []Education, experiences []Experience) {
	summary = &Summary{
		ID:       s.ID,
		AuthorID: s.Author.ID,
		Keywords: sql.NullString{String: s.Keywords},
	}

	for _, education := range s.Educations {
		educations = append(educations, Education{
			SummaryID:   summary.ID,
			Institution: sql.NullString{String: education.Institution},
			Speciality:  sql.NullString{String: education.Speciality},
			Graduated:   sql.NullString{String: education.Graduated},
			Type:        sql.NullString{String: education.Type},
		})
	}

	for _, experience := range s.Experiences {
		experiences = append(experiences, Experience{
			SummaryID:        summary.ID,
			CompanyName:      sql.NullString{String: experience.CompanyName},
			Role:             sql.NullString{String: experience.Role},
			Responsibilities: sql.NullString{String: experience.Responsibilities},
			Start:            sql.NullString{String: experience.Start},
			Stop:             sql.NullString{String: experience.Stop},
		})
	}

	return summary, educations, experiences
}

func toModel(s *Summary, eds []Education, exs []Experience, u *User, p *Person) *models.Summary {
	var educations []models.Education

	for _, ed := range eds {
		educations = append(educations, models.Education{
			Institution: ed.Institution.String,
			Speciality:  ed.Speciality.String,
			Graduated:   ed.Graduated.String,
			Type:        ed.Type.String,
		})
	}

	var experiences []models.Experience

	for _, ex := range exs {
		experiences = append(experiences, models.Experience{
			CompanyName:      ex.CompanyName.String,
			Role:             ex.Role.String,
			Responsibilities: ex.Responsibilities.String,
			Start:            ex.Start.String,
			Stop:             ex.Stop.String,
		})
	}

	nameArr := strings.Split(p.Name.String, " ")
	firstName := nameArr[0]
	var lastName string
	if len(nameArr) > 1 {
		lastName = nameArr[1]
	}

	birthday := fmt.Sprintf("%d-%02d-%02dT%02d:%02d:%02d-00:00\n", u.Registered.Time.Year(), u.Registered.Time.Month(),
		u.Registered.Time.Day(), u.Registered.Time.Hour(), u.Registered.Time.Minute(), u.Registered.Time.Second())

	author := models.Author{
		ID:        u.ID,
		Tag:       u.Tag.String,
		Email:     u.Email.String,
		Phone:     u.Phone.String,
		Avatar:    u.Avatar.String,
		FirstName: firstName,
		LastName:  lastName,
		Gender:    p.Gender.String,
		Birthday:  birthday,
	}

	return &models.Summary{
		ID:          s.ID,
		Author:      author,
		Keywords:    s.Keywords.String,
		Educations:  educations,
		Experiences: experiences,
	}
}

type GetOptions struct {
	userID uint64
}

type SummaryRepository struct {
	db *sql.DB
}

func NewSummaryRepository(db *sql.DB) *SummaryRepository {
	return &SummaryRepository{db}
}

func (r *SummaryRepository) CreateSummary(summary *models.Summary) (summaryID uint64, err error) {
	summaryDB, educationDBs, experienceDBs := toPostgres(summary)

	createSummary := `INSERT INTO summary (author, keywords)
					  VALUES ($1, $2) RETURNING id;`
	err = r.db.QueryRow(createSummary, summaryDB.AuthorID, summaryDB.Keywords).Scan(&summaryDB.ID)
	if err != nil {
		return summaryID, err
	}

	createEducation := `INSERT INTO education (summary_id, institution, speciality, graduated, type)
						VALUES ($1, $2, $3, $4, $5)`

	for _, educationDB := range educationDBs {
		_, err = r.db.Exec(createEducation, summaryDB.ID, educationDB.Institution, educationDB.Speciality,
						   educationDB.Graduated, educationDB.Type)
		if err != nil {
			return summaryID, err
		}
	}

	createExperience := `INSERT INTO experience (summary_id, company_name, role, responsibilities, start, stop)
						 VALUES ($1, $2, $3, $4, $5, $6)`

	for _, experienceDB := range experienceDBs {
		_, err = r.db.Exec(createExperience, summaryDB.ID, experienceDB.CompanyName, experienceDB.Role,
						   experienceDB.Responsibilities, experienceDB.Start, experienceDB.Stop)
		if err != nil {
			return summaryID, err
		}
	}

	return summaryDB.ID, nil
}

func (r *SummaryRepository) GetEducationsBySummaryID(summaryID uint64) ([]Education, error) {
	getEducations := `SELECT institution, speciality, graduated, type
					  FROM education WHERE summary_id = $1`

	rows, err := r.db.Query(getEducations, summaryID)
	if err != nil {
		return nil, err
	}

	var educationDBs []Education

	for rows.Next() {
		educationDB := Education{SummaryID: summaryID}

		err = rows.Scan(&educationDB.Institution, &educationDB.Speciality, &educationDB.Graduated,
			&educationDB.Type)
		if err != nil {
			return nil, err
		}

		educationDBs = append(educationDBs, educationDB)
	}

	return educationDBs, nil
}

func (r *SummaryRepository) GetExperiencesBySummaryID(summaryID uint64) ([]Experience, error) {
	getExperience := `SELECT company_name, role, responsibilities, start, stop
					  FROM experience WHERE summary_id = $1`

	rows, err := r.db.Query(getExperience, summaryID)
	if err != nil {
		return nil, err
	}

	var experienceDBs []Experience

	for rows.Next() {
		experienceDB := Experience{SummaryID: summaryID}

		err = rows.Scan(&experienceDB.CompanyName, &experienceDB.Role, &experienceDB.Responsibilities,
			&experienceDB.Start, &experienceDB.Stop)
		if err != nil {
			return nil, err
		}

		experienceDBs = append(experienceDBs, experienceDB)
	}

	return experienceDBs, nil
}

func (r *SummaryRepository) GetSummaryAuthor(authorID uint64) (*User, *Person, error) {
	user := User{ID: authorID}

	getUser := `SELECT person_id, tag, email, phone, avatar
				FROM users WHERE id = $1`
	err := r.db.QueryRow(getUser, user.ID).Scan(&user.PersonID, &user.Tag, &user.Email, &user.Phone, &user.Avatar)
	if err != nil {
		return nil, nil, err
	}

	var person Person

	getPerson := `SELECT name, gender, birthday
				  FROM person WHERE id = $1`
	err = r.db.QueryRow(getPerson, user.PersonID).Scan(&person.Name, &person.Gender, &person.Birthday)
	if err != nil {
		return nil, nil ,err
	}

	return &user, &person, nil
}

func (r *SummaryRepository) GetSummaries(opt *GetOptions) ([]models.Summary, error) {
	var rows *sql.Rows
	var err error

	if opt.userID == 0 {
		getSummaries := `SELECT id, author, keywords
					 	 FROM summary;`
		rows, err = r.db.Query(getSummaries)
		if err != nil {
			return nil, err
		}
	} else {
		getSummaries := `SELECT id, author, keywords
					 	 FROM summary WHERE author = $1;`
		rows, err = r.db.Query(getSummaries, opt.userID)
		if err != nil {
			return nil, err
		}
	}

	var summaries []models.Summary

	for rows.Next() {
		var summaryDB Summary

		err = rows.Scan(&summaryDB.ID, &summaryDB.AuthorID, &summaryDB.Keywords)
		if err != nil {
			return nil, err
		}

		educationDBs, err := r.GetEducationsBySummaryID(summaryDB.ID)
		if err != nil {
			return nil, err
		}

		experienceDBs, err := r.GetExperiencesBySummaryID(summaryDB.ID)
		if err != nil {
			return nil, err
		}

		userDB, personDB, err := r.GetSummaryAuthor(summaryDB.AuthorID)
		if err != nil {
			return nil, err
		}

		summaries = append(summaries, *toModel(&summaryDB, educationDBs, experienceDBs, userDB, personDB))
	}

	return summaries, nil
}

func (r *SummaryRepository) GetAllSummaries() (summaries []models.Summary, err error) {
	return r.GetSummaries(&GetOptions{})
}

func (r *SummaryRepository) GetUserSummaries(userID uint64) (summaries []models.Summary, err error) {
	return r.GetSummaries(&GetOptions{userID})
}

func (r *SummaryRepository) GetSummary(summaryID uint64) (*models.Summary, error) {
	summaryDB := Summary{ID: summaryID}

	getSummary := `SELECT author, keywords
				   FROM summary WHERE id = $1`
	err := r.db.QueryRow(getSummary, summaryID).Scan(&summaryDB.AuthorID, &summaryDB.Keywords)
	if err != nil {
		return &models.Summary{}, err
	}

	educationDBs, err := r.GetEducationsBySummaryID(summaryDB.ID)
	if err != nil {
		return &models.Summary{}, err
	}

	experienceDBs, err := r.GetExperiencesBySummaryID(summaryDB.ID)
	if err != nil {
		return &models.Summary{}, err
	}

	userDB, personDB, err := r.GetSummaryAuthor(summaryDB.AuthorID)

	return toModel(&summaryDB, educationDBs, experienceDBs, userDB, personDB), nil
}

func (r *SummaryRepository) ChangeSummary(summary *models.Summary) (err error) {
	summaryDB, educationDBs, experienceDBs := toPostgres(summary)

	changeSummary := `UPDATE summary
					  SET author = $1, keywords = $2
					  WHERE id = $3`
	_, err = r.db.Exec(changeSummary, summaryDB.AuthorID, summaryDB.Keywords, summaryDB.ID)
	if err != nil {
		return err
	}

	changeEducation := `UPDATE education
						SET institution = $1, speciality = $2, graduated = $3, type = $4
						WHERE summary_id = $5`

	for _, educationDB := range educationDBs {
		_, err = r.db.Exec(changeEducation, &educationDB.Institution, &educationDB.Speciality, &educationDB.Graduated,
						   &educationDB.Type, educationDB.SummaryID)
		if err != nil {
			return err
		}
	}

	changeExperience := `UPDATE experience
						SET company_name = $1, role = $2, responsibilities = $3, start = $4, stop = $5
						WHERE summary_id = $6`

	for _, experienceDB := range experienceDBs {
		_, err = r.db.Exec(changeExperience, &experienceDB.CompanyName, &experienceDB.Role,
						   &experienceDB.Responsibilities, &experienceDB.Start, &experienceDB.Stop,
						   &experienceDB.SummaryID)
		if err != nil {
			return err
		}
	}

	return nil
}

func (r *SummaryRepository) DeleteSummary(summaryID uint64) (err error) {
	deleteSummary := `DELETE FROM summary
					  WHERE id = $1`
	_, err = r.db.Exec(deleteSummary, summaryID)
	if err != nil {
		return err
	}

	deleteEducations := `DELETE FROM education
						 WHERE summary_id = $1`
	_, err = r.db.Exec(deleteEducations, summaryID)
	if err != nil {
		return err
	}

	deleteExperiences := `DELETE FROM experience
						  WHERE summary_id = $1`
	_, err = r.db.Exec(deleteExperiences, summaryID)
	if err != nil {
		return err
	}

	return nil
}