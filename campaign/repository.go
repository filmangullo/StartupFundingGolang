package campaign

type Repository struct {
    FindAll() ([]Campaign, error)
    FindByUserID(userID int) ([]Campaign, error)
}

type repository struct {
    db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
    return &repository{db}
}

func FindAll(r *repository) ([]Campaign, error) {
    var campaigns []Campaign
    err := r.db.find(&campaigns).Error
    if err != nil {
        return campaigns, err
    }

    return campaigns, nil
}

func (r *repository) FindByUserID(userID int) ([]Campaign, error) {
    var campaigns []Campaign
    err:= r.db.Where("user_id=?", userID).Find(&campaigns).Error

    if err != nil {
        return campaigns, err
    }
    return campaigns, nil
}