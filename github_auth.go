package main

// func githubAuthorizationHandler(w http.ResponseWriter, r *http.Request) {
// 	url := oauthGitConfig.AuthCodeURL("")
// 	http.Redirect(w, r, url, http.StatusFound)
// }
//
// func gitHomeHandler(w http.ResponseWriter, r *http.Request) {
// 	code := r.FormValue("code")
// 	token, err := oauthGitConfig.Exchange(oauth2.NoContext, code)
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 		return
// 	}
//
// 	client := oauthGitConfig.Client(oauth2.NoContext, token)
// 	response, _ := client.Get(githubProfileInfoURL)
//
// 	defer response.Body.Close()
// 	rawBody, err := ioutil.ReadAll(response.Body)
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 		return
// 	}
//
// 	s, _ := session.Manager.SessionStart(w, r)
// 	defer s.SessionRelease(w)
//
// 	s.Set("id_token", token.Extra("id_token"))
// 	s.Set("access_token", token.AccessToken)
// 	s.Set("profile", string(rawBody))
//
// 	http.Redirect(w, r, "/user", http.StatusMovedPermanently)
// }
