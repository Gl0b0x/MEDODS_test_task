package handler

import (
	"MEDODS/pkg/models"
	"MEDODS/pkg/service"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

func (h *Handler) refresh(w http.ResponseWriter, r *http.Request) {
	var req models.RefreshRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		h.logger.Println("invalid request body", err.Error())
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()
	if req.RefreshToken == "" {
		h.logger.Println("refresh token is required")
		http.Error(w, "Refresh token is required", http.StatusUnauthorized)
		return
	}
	refreshToken, err := service.DecodeBase64(req.RefreshToken)
	if err != nil {
		h.logger.Println("invalid refresh token:", err.Error())
		http.Error(w, "Invalid refresh token", http.StatusUnauthorized)
	}
	uuid, err := h.services.Authorization.ParseToken(refreshToken)
	if err != nil {
		h.logger.Println("invalid refresh token:", err.Error())
		http.Error(w, "Invalid refresh token", http.StatusUnauthorized)
		return
	}
	user, err := h.services.Authorization.GetUserByUUID(uuid)
	if err != nil {
		h.logger.Println("user not found:", err.Error())
		http.Error(w, "User not found", http.StatusUnauthorized)
		return
	}
	err = h.services.Authorization.CompareHashAndPassword(*user.RefreshToken, req.RefreshToken)
	if err != nil {
		h.logger.Println("invalid refresh token:", err.Error())
		http.Error(w, "Invalid refresh token", http.StatusUnauthorized)
		return
	}
	if user.Ip == nil {
		*user.Ip = r.RemoteAddr
	}
	if user.Ip != nil && *user.Ip != r.RemoteAddr {
		h.logger.Println("Отправка оповещения на почту...")
	}
	accessToken, refreshToken, err := h.services.GenerateTokens(user)
	if err != nil {
		h.logger.Println(err.Error())
		http.Error(w, "Internal server error", http.StatusInternalServerError)
	}
	base64RefreshToken := service.EncodeBase64(refreshToken)
	err = h.services.UpdateUserRefreshToken(user, base64RefreshToken)
	if err != nil {
		h.logger.Println(err.Error())
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	ans := models.LoginResponse{AccessToken: accessToken, RefreshToken: base64RefreshToken}
	if err = json.NewEncoder(w).Encode(ans); err != nil {
		h.logger.Println(err.Error())
		http.Error(w, "Failed to encode token data", http.StatusInternalServerError)
		return
	}
}

//func (h *Handler) authMiddleware(next http.Handler) http.Handler {
//	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
//		authHeader := strings.Split(r.Header.Get("Authorization"), " ")
//		if len(authHeader) != 2 && authHeader[0] != "Bearer" {
//			http.Error(w, "Invalid auth header", http.StatusUnauthorized)
//			return
//		}
//		if len(authHeader[1]) == 0 {
//			http.Error(w, "token is empty", http.StatusUnauthorized)
//		}
//		uuid, err := h.services.Authorization.ParseToken(authHeader[1])
//		if err != nil {
//			http.Error(w, "Invalid token", http.StatusUnauthorized)
//			return
//		}
//		ctx := context.WithValue(r.Context(), "uuid", uuid)
//		next.ServeHTTP(w, r.WithContext(ctx))
//	})
//}

func (h *Handler) login(w http.ResponseWriter, r *http.Request) {
	uuid := r.FormValue("uuid")
	if uuid == "" {
		h.logger.Println("Invalid User guid")
		http.Error(w, "Invalid User guid", http.StatusBadRequest)
		return
	}
	user, err := h.services.Authorization.GetUserByUUID(uuid)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			h.logger.Println("User with uuid " + uuid + " not found")
			http.Error(w, "User not found", http.StatusNotFound)
			return
		}
		h.logger.Println(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	ip := r.RemoteAddr
	if user.Ip == nil {
		user.Ip = &ip
	}
	if ip != *user.Ip {
		fmt.Println("Отправка оповещения на почту")
	}
	accessToken, refreshToken, err := h.services.GenerateTokens(user)
	if err != nil {
		h.logger.Println(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	base64RefreshToken := service.EncodeBase64(refreshToken)
	err = h.services.UpdateUserRefreshToken(user, base64RefreshToken)
	if err != nil {
		h.logger.Println(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	ans := models.LoginResponse{AccessToken: accessToken, RefreshToken: base64RefreshToken}
	if err = json.NewEncoder(w).Encode(ans); err != nil {
		h.logger.Println("Failed to encode token data", err.Error())
		http.Error(w, "Failed to encode token data", http.StatusInternalServerError)
		return
	}
}
