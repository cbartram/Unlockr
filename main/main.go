package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const OAUTH_URL = "https://nestauthproxyservice-pa.googleapis.com/v1/issue_jwt"
const API_KEY = "<redacted>"
const BEARER_TOKEN = "<redacted>"

func main() {
	client := &http.Client{}
	fmt.Println("Retrieving OAuth token")
	getAccessToken(*client)
}

type OAuthResponse struct {
	expire_after                    string
	policy_id                       string
	google_oauth_access_token       string
	embed_google_oauth_access_token string
}

func getAccessToken(client http.Client) OAuthResponse {
	var jsonResponse OAuthResponse
	var jsonBody = []byte(`{ "expire_after":"3600s", "policy_id":"authproxy-oauth-policy", "google_oauth_access_token":` + BEARER_TOKEN + `, "embed_google_oauth_access_token":"true" }`)
	fmt.Println("Making Post Request to: " + OAUTH_URL)

	req, _ := http.NewRequest("POST", OAUTH_URL, bytes.NewBuffer(jsonBody))
	req.Header.Add("x-goog-api-key", API_KEY)
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Host", "nestauthproxyservice-pa.googleapis.com")
	req.Header.Add("Authorization", "Bearer "+BEARER_TOKEN)
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	defer res.Body.Close()

	bodyBytes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
	}
	bodyString := string(bodyBytes)
	fmt.Println("Response Body: " + bodyString)

	if res.StatusCode == http.StatusOK {
		err := json.NewDecoder(res.Body).Decode(&jsonResponse)
		if err != nil {
			fmt.Println(jsonResponse)
			return jsonResponse
		} else {
			fmt.Println(err)
		}
	}
	return OAuthResponse{}
}

//{
//	"expire_after":"3600s",
//	"policy_id":"authproxy-oauth-policy",
//	"google_oauth_access_token":"ya29.a0Ae4lvC0p4usjMdR4C1nJ2FeX1t9J856PZGLU_8GlMxb4Aa4c4cnH5Gta0vtaLym17acMsFDLdFo5Is_KZhsSnlLjoQrJ23txMpsv6LUgQpV31WWLoPqg7Xbgmb81mWAFEwk1wa0-02TEEdb0TfEA92PZhslYWy1qCrI",
//	"embed_google_oauth_access_token":"true"
//}

//
///**
// * Retrieves a list of recent events that the Nest camera detected. It can take two optional params
// * start and end which are unix timestamps in seconds since epoch and represent a window of time to retrieve
// * events for.
// * @returns {Promise<any>}
// */
//const getEvents = async (start = null, end = null) => {
//const options = {
//'method': 'GET',
//'url': 'https://nexusapi-us1.dropcam.com/cuepoint/92b207f519f248409f09e60dfc6853e0/2',
//'headers': {
//'Authorization': `Basic ${access_token}`
//}
//};
//try {
//return JSON.parse(await request(options));
//} catch(e) {
//console.log('[ERROR] Failed to retrieve events from the Nest API: ', e)
//}
//};
//
//getOAuthToken().then(({ jwt }) => {
//console.log('[INFO] Got OAuth Access token: ', jwt);
//getEvents(jwt).then(events => {
//console.log('[INFO] Retrieved: ', events, " events");
//console.log(events.map(event => event.id));
//});
//});
