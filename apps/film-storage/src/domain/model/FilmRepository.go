package model

import (
	"context"
	"errors"
	"film-storage/adapter"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	idKeyName      string = "_id"
	collectionName string = "movies"
)

type FilmRepository struct {
	db         *adapter.DBClient
	collection *mongo.Collection
}

func NewRepository(db *adapter.DBClient) *FilmRepository {
	collection := db.Collection(collectionName)
	return &FilmRepository{db: db, collection: collection}
}

func (dao *FilmRepository) HealthCheck() error {
	return dao.db.Ping()
}

func (dao *FilmRepository) DeleteById(idFilm string) error {
	ctx, cancel := initContext()
	defer cancel()
	id, err := primitive.ObjectIDFromHex(idFilm)
	if err != nil {
		return err
	}
	_, err = dao.collection.DeleteOne(ctx, bson.M{idKeyName: id})
	if err != nil {
		return err
	}
	return nil
}

func (dao *FilmRepository) FindById(idFilm string) (*Film, error) {
	ctx, cancel := initContext()
	defer cancel()
	id, err := primitive.ObjectIDFromHex(idFilm)
	if err != nil {
		return nil, err
	}
	cur := dao.collection.FindOne(ctx, bson.M{idKeyName: id})
	if cur.Err() != nil {
		log.Printf("Error: %v", cur.Err())
		return nil, cur.Err()
	}
	var f Film
	err = cur.Decode(&f)
	if err != nil {
		return nil, err
	}
	return &f, nil
}

func (dao *FilmRepository) GetAll() ([]Film, error) {
	ctx, cancel := initContext()
	defer cancel()
	cursor, err := dao.collection.Find(ctx, bson.M{})

	if err != nil {
		return []Film{}, err
	}

	filmList := []Film{}
	for cursor.Next(ctx) {
		var f Film
		err = cursor.Decode(&f)
		filmList = append(filmList, f)
	}

	return filmList, nil
}

func (dao *FilmRepository) CreateOne(film Film) (Film, error) {
	ctx, cancel := initContext()
	defer cancel()
	r, err := dao.collection.InsertOne(ctx, film)
	if err != nil || r.InsertedID == nil {
		if err == nil {
			return film, errors.New("No object id")
		}
		return film, err
	}
	return film, nil
}

func initContext() (context.Context, context.CancelFunc) {
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	return ctx, cancel
}

/*
func (dao *FilmRepository) InsertOne(film Film) error {
	dao.db.
}

func (dao *FilmRepository) FindAllByType(symbol string) []Candle {
	query := dao.db.Query(CollectionName).Where(c.Field("symbol").Eq(symbol))
	docs := dao.db.FindAllByCriteria(query)
	var candle *Candle
	var candles []Candle = []Candle{}
	for _, doc := range docs {
		doc.Unmarshal(&candle)
		candles = append(candles, *candle)
		log.Println(candle)
	}
	return candles
}

func (dao *FilmRepository) FindCandle(symbol string, market string, precision string, ts int64) (*Candle, error) {
	query := dao.db.Query(CollectionName).
		Where((*c.Criteria)(c.Field("symbol").Eq(symbol).
			And((*c.Criteria)(c.Field("market").Eq(market))).
			And((*c.Criteria)(c.Field("precision").Eq(precision))).
			And((*c.Criteria)(c.Field("ts").Eq(ts)))))
	docs := dao.db.FindAllByCriteria(query)
	if len(docs) > 1 {
		return nil, fmt.Errorf("%s (%d)", "no unique result", len(docs))
	} else if len(docs) == 0 {
		return nil, nil
	} else {
		var candle *Candle
		docs[0].Unmarshal(&candle)
		return candle, nil
	}
}

func (dao *FilmRepository) Save(candle *Candle) error {
	c, err := dao.FindCandle(candle.Symbol, candle.Market, candle.Precision, candle.Ts)
	if err != nil {
		log.Printf("error saving a document %s", err)
	} else {
		if c != nil {
			candle.Id = c.Id
			candle.Version = c.Version + 1
			return dao.UpdateOne(candle)
		} else {
			data := convert(candle)
			return dao.db.InsertOne(data, CollectionName)
		}
	}
	return nil
}

func (dao *FilmRepository) UpdateOne(candle *Candle) error {
	return dao.db.Query(CollectionName).
		Where((*c.Criteria)(c.Field("_id").Eq(candle.Id))).Update(convert(candle))
} */
