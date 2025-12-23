package universities

import (
	"database/sql"
	"strconv"

	"github.com/gin-gonic/gin"

	"backend/internal/response"
)

type Handler struct {
	svc *Service
}

func NewHandler(svc *Service) *Handler {
	return &Handler{svc: svc}
}

func (h *Handler) Search(c *gin.Context) {
	countryCode := c.Query("country_code")
	q := c.Query("q")

	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	size, _ := strconv.Atoi(c.DefaultQuery("size", "20"))

	items, total, err := h.svc.Search(c.Request.Context(), SearchParams{
		CountryCode: countryCode,
		Q:           q,
		Page:        page,
		Size:        size,
	})
	if err != nil {
		response.ServerError(c, err.Error())
		return
	}

	dtoItems := make([]UniversityListItemDTO, 0, len(items))
	for _, u := range items {
		dtoItems = append(dtoItems, UniversityListItemDTO{
			ID:          u.ID,
			CountryCode: u.CountryCode,
			Country:     u.Country,
			NameEN:      u.NameEN,
			NameENShort: u.NameENShort,
			NameCN:      u.NameCN,
		})
	}

	response.OK(c, PagedResult[UniversityListItemDTO]{
		Page:  page,
		Size:  size,
		Total: total,
		Items: dtoItems,
	})
}

func (h *Handler) GetByID(c *gin.Context) {
	idStr := c.Param("id")
	id64, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		response.BadRequest(c, "invalid id")
		return
	}

	u, err := h.svc.GetByID(c.Request.Context(), id64)
	if err != nil {
		if err == sql.ErrNoRows {
			response.NotFound(c, "university not found")
			return
		}
		response.ServerError(c, err.Error())
		return
	}

	// 详情先简单返回（你后面可把 domains_json 解析成 []string）
	response.OK(c, gin.H{
		"id":            u.ID,
		"country_code":  u.CountryCode,
		"country":       u.Country,
		"name_en":       u.NameEN,
		"name_en_short": u.NameENShort,
		"name_cn":       u.NameCN,
		"domains_json":  u.DomainsJSON,
	})
}


func (h *Handler) ListAllNameCN(c *gin.Context) {
	names, err := h.svc.ListAllNameCN(c.Request.Context())
	if err != nil {
		response.ServerError(c, err.Error())
		return
	}
	response.OK(c, names)
}

func (h *Handler) OptionsCN(c *gin.Context) {
	q := c.Query("q")
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	size, _ := strconv.Atoi(c.DefaultQuery("size", "20"))

	items, total, err := h.svc.OptionsCN(c.Request.Context(), OptionsCNParams{
		Q:    q,
		Page: page,
		Size: size,
	})
	if err != nil {
		response.ServerError(c, err.Error())
		return
	}

	response.OK(c, PagedResult[UniversityOptionCNDTO]{
		Page:  page,
		Size:  size,
		Total: total,
		Items: items,
	})
}
