package internal

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"

	"github.com/gorilla/mux"

	"./link"
	"./netutil"
	v "./validation"
)

type Handler struct {
	linkService link.Service
}

func NewHandler(linkService link.Service) *Handler {
	return &Handler{
		linkService: linkService,
	}
}

func (h *Handler) GetHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	userID := params["id"]
	if userID == "" {
		netutil.WriteResponse(errors.New("userID cannot be empty"), http.StatusBadRequest, w)
		return
	}

	sortByDateCreated, err := h.extractQueryParams(r.URL.Query())
	if err != nil {
		netutil.WriteResponse(err.Error(), http.StatusBadRequest, w)
	}
	linkData, err := h.linkService.FetchLinksForUser(userID, sortByDateCreated)
	if err != nil {
		errorMsg := fmt.Sprintf("{\"message\": \"A system error occurred. Details: %+v \"} ", err)
		netutil.WriteResponse(errorMsg, http.StatusInternalServerError, w)
	}
	netutil.WriteResponse(linkData, http.StatusOK, w)
}

func (h *Handler) PostHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	userID := params["id"]
	if userID == "" {
		netutil.WriteResponse(errors.New("userID cannot be empty"), http.StatusBadRequest, w)
		return
	}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		errorMsg := fmt.Sprintf("{\"message\": \"A system error occurred. Details: %+v \"} ", err)
		netutil.WriteResponse(errorMsg, http.StatusInternalServerError, w)
		return
	}
	var createLinkRequest *link.CreateLinkRequest
	createLinkRequest, err = h.validateCreateLinkRequestBody(body)
	if err != nil {
		netutil.WriteResponse(err.Error(), http.StatusBadRequest, w)
		return
	}
	linkData, err := h.linkService.Create(userID, createLinkRequest)
	if err != nil {
		errorMsg := fmt.Sprintf("{\"message\": \"A system error occurred. Details: %+v \"} ", err)
		netutil.WriteResponse(errorMsg, http.StatusInternalServerError, w)
	}
	netutil.WriteResponse(linkData, http.StatusCreated, w)
}

func (h *Handler) extractQueryParams(queryParams url.Values) (bool, error) {
	sortByDateCreated := false
	var err error
	sortByDateCreatedQueryParam := queryParams.Get("sortByDateCreated")
	if sortByDateCreatedQueryParam != "" {
		sortByDateCreated, err = strconv.ParseBool(sortByDateCreatedQueryParam)
		if err != nil {
			err = errors.New("{\"message\": \"sortByDateCreated has to be a boolean value true or false. \"}")
		}
	}
	return sortByDateCreated, err
}

// TODO This needs to be somehow generalised so that the switch statement doesn't grow exponentially when a new link type is added
func (h *Handler) validateCreateLinkRequestBody(createLinkRequestBody []byte) (*link.CreateLinkRequest, error) {
	var data json.RawMessage
	createLinkRequest := link.CreateLinkRequest{Data: &data}
	if err := json.Unmarshal(createLinkRequestBody, &createLinkRequest); err != nil {
		return nil, err
	}
	switch createLinkRequest.Type {
	case "classic":
		{
			var classicLink link.Classic
			if err := json.Unmarshal(data, &classicLink); err != nil {
				return nil, err
			}
			err := v.ValidateClassicLink(classicLink)
			if err != nil {
				return nil, err
			}
		}
	case "shows_list":
		{
			var showsList link.ShowsList
			if err := json.Unmarshal(createLinkRequestBody, &showsList); err != nil {
				return nil, err
			}
			err := v.ValidateShowsList(showsList)
			if err != nil {
				return nil, err
			}
		}
	case "music_player":
		{
			var musicPlayer link.MusicPlayer
			if err := json.Unmarshal(createLinkRequestBody, &musicPlayer); err != nil {
				return nil, err
			}
			err := v.ValidateMusicPlayer(musicPlayer)
			if err != nil {
				return nil, err
			}
		}
	default:
		return nil, errors.New("invalid link type")
	}

	return &createLinkRequest, nil
}
