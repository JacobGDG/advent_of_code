package getinput

import (
  "io/ioutil"
  "net/http"
  "strconv"
)

func handleError(err error) {
  if err != nil {
    panic(err)
  }
}

func request(token string, day int) *http.Request {
  const base_input_url, final_route string = "https://adventofcode.com/2020/day/", "/input"
  var full_url string = base_input_url + strconv.Itoa(day) + final_route

  request, err := http.NewRequest("GET", full_url, nil)

  handleError(err)

  request.Header.Add("Cookie", "session=" + token)

  return request
}

func GetInput(token string) string {
  client := &http.Client{}
  var day int = 3

  resp, err := client.Do(request(token, day))

  handleError(err)

  defer resp.Body.Close()

  body, err := ioutil.ReadAll(resp.Body)

  handleError(err)

  return string(body)
}
