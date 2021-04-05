<template>
  <div class="main">
    <div class="left-panel">
      <div class="player-container">
        <md-switch v-if="!state.Started" v-model="ready" @change="changeReady"
          >Are you ready?</md-switch
        >
      </div>
      <div
        v-for="(element, index) in state.names"
        :key="index"
        class="player-container"
      >
        <div v-if="element !== ''">
          <img
            src="https://images-na.ssl-images-amazon.com/images/I/81RfjOhvMrL.png"
            alt="Ikona gracza"
            class="player-image"
          />
          {{ element + (state.players_status[index] == 1 ? " Ready!" : "") }}
        </div>
      </div>
    </div>
    <div v-show="false">
      <img
        src="https://upload.wikimedia.org/wikipedia/commons/thumb/9/91/Menschenaergern.svg/1200px-Menschenaergern.svg.png"
        width="1200"
        height="1200"
        id="ludo-img"
      />
    </div>
    <div class="game">
      <div
        v-for="(position, index) in positions"
        :key="index"
        v-bind:style="{
          position: 'absolute',
          top: position.y - 22 + 'px',
          left: position.x - 22 + 'px',
          width: '45px',
          height: '45px',
          zIndex: '3',
        }"
        @click="choosePawn(field_pawn_number[index])"
        :class="
          field_pawn_number[index] !== -1 && playerStatus === CHOOSING_PAWN
            ? 'field-has-players-pawn'
            : ''
        "
        :id="'field' + index"
      >
        {{ index }}
      </div>
      <canvas id="game" ref="canvas" width="800" height="800"> </canvas>
    </div>

    <div v-if="id === state.turn">
      <md-button
        @click.once="nextStatus"
        :key="diceKey"
        style="pointer: move"
        v-if="playerStatus === ROLLING_DICE"
        >Rzuć Kostka!</md-button
      >
      <md-progress-spinner
        v-if="playerStatus === DICE_ROLLED"
        md-mode="indeterminate"
        :md-diameter="30"
      ></md-progress-spinner>
    </div>

    <div v-else-if="state.started">
      <h5>{{ state.roll }}</h5>
    </div>
  </div>
</template>

<script lang="ts">
import { Component, Vue, Ref } from "vue-property-decorator";
import Cookies from "js-cookie";
import * as api from "@/store/api";
import { GameState, board_positions, player_colors } from "@/types/index";
@Component({
  components: {},
})
export default class Game extends Vue {
  NOT_READY = 0;
  READY = 1;
  ROLLING_DICE = 2;
  DICE_ROLLED = 3;
  CHOOSING_PAWN = 4;
  PAWN_CHOSEN = 5;
  chosen = 0;

  canvas?: HTMLCanvasElement;
  canvasCtx: CanvasRenderingContext2D;
  state: GameState = new GameState();
  positions: Array<any> = [];
  field_has_pawn: Array<boolean> = [];
  field_pawn_number: Array<number> = [];

  ready = false;
  id = -10;
  type = "";
  key = "";
  player_name = "";
  game_id = "-1";
  playerStatus = 0;
  image: any;

  diceKey = 0;

  nextStatus() {
    this.playerStatus++;
  }

  choosePawn(pawn: number) {
    if (pawn === -1) return;
    if (this.playerStatus === this.CHOOSING_PAWN) {
      this.nextStatus();
      this.chosen = pawn;
    }
  }

  async fetchState() {
    const player_status = this.playerStatus.toString();
    const chosen = this.chosen.toString();
    const key = this.key;
    const game_id = this.game_id;
    console.log("sending player status", this.playerStatus);
    const response = await api.fetchState({
      key,
      chosen,
      player_status,
      game_id,
    });
    this.state = response.data;
    this.playerStatus = this.state.players_status[this.id];
    console.log("RECEIVED PLAYER STATUS", this.playerStatus);
    this.diceKey++;
  }

  changeReady() {
    const pi = this.state.names.findIndex((el) => el === this.player_name);
    console.log(this.state, pi);
    if (pi !== -1) {
      this.state.players_status[pi] =
        this.state.players_status[pi] == 0 ? 1 : 0;
    }
    this.playerStatus = this.state.players_status[pi];
  }

  mounted() {
    this.field_has_pawn = [];
    board_positions.forEach((el) => {
      this.field_pawn_number.push(-1);
    });

    this.positions = board_positions;
    this.player_name = Cookies.get("player_name");
    this.key = Cookies.get("key");
    this.game_id = Cookies.get("game_id");
    this.id = parseInt(Cookies.get("player_number"));

    this.canvas = document.getElementById("game") as HTMLCanvasElement;
    this.canvasCtx = this.canvas.getContext("2d") as CanvasRenderingContext2D;
    this.image = document.getElementById("ludo-img");

    setInterval(this.drawImage, 20);
    setInterval(this.fetchState, 3000);
  }

  drawImage() {
    this.canvasCtx?.drawImage(this.image, 0, 0, 800, 800);
    this.field_pawn_number.map(() => {
      return -1;
    });
    this.state.positions.forEach((player, index) => {
      player.forEach((pawn, pawnPlayerNumber) => {
        const position = board_positions[pawn];
        if (index == this.id) {
          this.field_pawn_number[pawn] = pawnPlayerNumber;
        }
        this.canvasCtx?.beginPath();
        this.canvasCtx.fillStyle = player_colors[index];
        this.canvasCtx.arc(position.x, position.y, 20, 0, 2 * Math.PI);
        this.canvasCtx.fill();
        this.canvasCtx.stroke();
      });
    });
  }
}
</script>

<style scoped>
.main {
  display: flex;
  flex-direction: row;
  height: 100%;
}
.game {
  max-width: 800px;
  width: 800px;
  height: 800px;
  flex: 1;
  z-index: 1;
  position: relative;
}
.player-image {
  border-radius: 100%;
  width: 120px;
  height: 120px;
}
.player-container {
  margin-bottom: 2rem;
}
.left-panel {
  width: 20%;
  height: 100%;
  background-color: #bdbdbd;
  display: flex;
  flex-direction: column;
  align-content: stretch;
}
.field-has-players-pawn {
  cursor: pointer;
  z-index: 2;
}
.field-has-players-pawn:hover {
  background-color: rebeccapurple;
  opacity: 0.5;
  border-radius: 100%;
}
</style>
