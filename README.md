# Ad

```JSON
{
  "id": "string",
  "name": "string",
  "filters": {
    "type": "string",
    "h": "integer",
    "w": "integer"
  },
  "price": "float"
}
```

---

# Bid Request

```JSON
{
  "id": "string",
  "imp": [{
    "id": "string",
    "banner": {
      "w": "integer",
      "h": "integer"
    },
    "video": {
      "w": "integer",
      "h": "integer"
    },
    "bidfloor": "float"
  }],
  "device": {
    "geo": {
      "lat": "float",
      "lon": "float",
      "country": "string"
    },
    "deviceType": "integer",
    "ip": "string",
    "ifa": "string"
  },
  "cur": "string"
}
```

---

# Bid Response
```JSON
{
  "id": "string",
  "seatbid": [{
    "seat": "string",
    "bid": {
      "id": "string",
      "impid": "string",
      "price": "float",
      "adm": "string",
      "adid" :"string"
    }
  }],
  "bidid": "string"

}
```

# Routes
| METHOD | Path | Description |
|---|---|---|
| GET | `/health` | returns 200 |
| GET | `/bidrequest` | returns 200 + Bid Response or 204 | 
| GET | `/ads` | returns all ads |
| GET | `/ads/:id` | returns ad with id or 204 |
| POST | `/ads` | creates an ad |
| PUT | `/ads/:id` | updates an ad |
| DELETE | `/ads/:id` | deletes an ad |
| GET | `/tree` | returns whole tree of sorted ads. Ads are sorted so it's faster to find them at bidding time. |
| GET | `/tree/:key` | returns the leaf for given key. (A leaf is a map of Ads that respect same)|

# TODO
 - [x] Respect golang standard directory structure https://github.com/golang-standards/project-layout
 - [ ] Add swagger
 - [ ] Implements `/bidrequest` route
 - [ ] Add unit tests
 - [ ] Add system tests