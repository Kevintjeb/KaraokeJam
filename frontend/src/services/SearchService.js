import axios from "axios";
import { searchEndpoint } from "../constants";
export class SearchService {
  search(query) {
    return axios.get(`${searchEndpoint}/${query}`);
  }
}
