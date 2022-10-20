package model

import (
	"film-storage/adapter"
	"fmt"
	"log"
)

type FilmRepository struct {
	db *adapter.DBClient
}

func NewRepository(db *adapter.DBClient) *FilmRepository {
	return &FilmRepository{db: db}
}

func (dao *FilmRepository) InsertOne(film Film) error {

}

func (dao *CandleRepository) FindAllByType(symbol string) []Candle {
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

func (dao *CandleRepository) FindCandle(symbol string, market string, precision string, ts int64) (*Candle, error) {
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

func (dao *CandleRepository) Save(candle *Candle) error {
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

func (dao *CandleRepository) UpdateOne(candle *Candle) error {
	return dao.db.Query(CollectionName).
		Where((*c.Criteria)(c.Field("_id").Eq(candle.Id))).Update(convert(candle))
}
