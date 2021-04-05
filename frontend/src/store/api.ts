import axios, { AxiosResponse } from "axios";
import * as consts from "@/consts";

export const api = axios.create({
  baseURL: consts.api_address + "api/",
});

export async function registerPlayer(name: string): Promise<AxiosResponse> {
  console.log("** register player...", consts.api_address);
  return await api.post("join", { name });
}

export async function fetchState(player_data: {
  key: string;
  player_status: string;
  chosen: string;
  game_id: string;
}) {
  //console.log("** fetching game state...", consts.api_address + "process");
  return await api.post("process", player_data);
}
