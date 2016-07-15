package nl

import (
    "time"
    "golang.org/x/net/context"
	"google.golang.org/appengine/datastore"
)

type Tag struct {
    Id string
    LastVisited time.Time
    Score   float64
}

func (tag *Tag) StoreTag(ctx context.Context, id string)(err error) {
    key := datastore.NewIncompleteKey(ctx, "Tag", nil)
    _, err = datastore.Put(ctx, key, tag)

    return
}

func (tag *Tag) GetTag(ctx context.Context, id string)(err error) {
    key := datastore.NewKey(ctx, "Tag", id, 0, nil)
    err = datastore.Get(ctx, key, tag)

    return
}

func (tag *Tag) CheckOffer()(offer string){
    offer = "false"
    
    if tag.Score > 0.0 {
        offer = "true"
    }
    return
}


