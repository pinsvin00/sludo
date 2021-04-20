<template>
  <div class="main">
    <div class="endscreen" v-if="isGameEnded">
      <div class="centered-box">
        <div>Wygrał gracz numer {{ winner }}</div>
      </div>
    </div>
    <div class="left-panel">
      <div class="player-container">
        <md-switch v-if="!state.Started" v-model="ready" @change="changeReady"
          >Gotowy?</md-switch
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
        @mouseover="highlightPosition(index)"
        @mouseleave="dehighlightPosition(index)"
        :class="{
          'field-has-players-pawn': field_pawn_number[index] !== -1,
          'possible-move': isAbleToMove[index],
        }"
        :id="'field-' + index"
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

    <img :src="diceImgUrl[state.roll]" class="dice" />
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
  /*
  TODO
  ->Konczenie rozgrywki
  ->Czat
  */

  NOT_READY = 0;
  READY = 1;
  ROLLING_DICE = 2;
  DICE_ROLLED = 3;
  CHOOSING_PAWN = 4;
  PAWN_CHOSEN = 5;
  chosen = 0;

  isGameEnded = false;
  winner = -1;

  canvas?: HTMLCanvasElement;
  canvasCtx: CanvasRenderingContext2D;
  state: GameState = new GameState();
  positions: Array<any> = [];
  field_has_pawn: Array<boolean> = [];
  field_pawn_number: Array<number> = [];
  isAbleToMove: Array<boolean> = [];

  ready = false;
  readOnce = false;
  id = -10;
  type = "";
  key = "";
  player_name = "";
  game_id = "-1";
  playerStatus = 0;
  image: any;

  diceImgUrl = [
    "https://cdn.pixabay.com/photo/2013/07/12/17/39/dice-152173_960_720.png",
    "https://cdn.pixabay.com/photo/2013/07/12/17/39/dice-152173_960_720.png",
    "https://cdn.pixabay.com/photo/2013/07/12/17/39/dice-152174_1280.png",
    "https://lh3.googleusercontent.com/proxy/TnzHE9DfCb0cZ7qfRSTeRoOHWF_Q_ReWBWBKYkF1PtLj4iEySmRm_UPuFmuyQrBh5PxnaO_CRtifX4GAKmfBnAh_NpZwm7MpxeASaB-mVtv2MqE",
    "https://freesvg.org/img/dado-4.png",
    "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcROGmw-bqWkz6Icg7_Kemdv-3DndU1Gn2-N5Sl-KsYonvHxGCWcdDD98_nH9ejih6QU8ik&usqp=CAU",
    "https://lh3.googleusercontent.com/proxy/xCIio31m93KAR8r5s9XdolqkJTwZzGm-z4NxDFvTCks8Y0GEpn1ewmEq8HQvNjNKUAGJULMcvafpOyb7242Ou8Ts1QA4WyDnT-4YwyKlYRSuM4c",
  ];

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

  highlightPosition(position: number) {
    if (this.field_pawn_number[position] <= -1 || this.state.turn !== this.id) {
      return;
    }
    const positionToHighlight = this.calculate_position(position);
    if (positionToHighlight === -1) return;
    const elementHighlight = document.getElementById(
      `field-${positionToHighlight}`
    );
    if (elementHighlight) {
      elementHighlight.className += ` pinkHover`;
    } else {
      console.log("ERR : NULL ELEMENT!");
    }
  }

  dehighlightPosition() {
    const highlighted = document.getElementsByClassName("pinkHover");
    for (let i = 0; i < highlighted.length; i++) {
      highlighted[i].classList.remove("pinkHover");
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
    if (this.state.players_status[this.id] === this.CHOOSING_PAWN) {
      if (!this.readOnce) {
        const synth = window.speechSynthesis;
        const speech = new SpeechSynthesisUtterance(this.state.roll.toString());
        synth.speak(speech);
        this.readOnce = true;
      }
    } else {
      this.readOnce = false;
    }
    this.isGameEnded = this.state.ended;
    this.winner = this.state.winner;
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
    this.isAbleToMove = [];
    board_positions.forEach((el) => {
      this.field_pawn_number.push(-1);
      this.isAbleToMove.push(false);
    });

    this.positions = board_positions;
    this.player_name = Cookies.get("player_name");
    this.key = Cookies.get("key");
    this.game_id = Cookies.get("game_id");
    this.id = parseInt(Cookies.get("player_number"));

    this.canvas = document.getElementById("game") as HTMLCanvasElement;
    this.canvasCtx = this.canvas.getContext("2d") as CanvasRenderingContext2D;
    this.image = document.getElementById("ludo-img");

    setInterval(this.drawImage, 300);
    setInterval(this.fetchState, 3000);
  }

  calculate_position(position: number) {
    console.log(position);
    if (position <= 16) {
      if (this.state.roll === 1 || this.state.roll === 6) {
        return 16 + this.id * 10;
      }
      return -1;
    } else {
      position = (position + this.state.roll) % (40 + 15);
      if (position <= 15) {
        position += 15;
      }
      return position;
    }
  }

  drawImage() {
    this.canvasCtx?.drawImage(this.image, 0, 0, 800, 800);
    this.field_pawn_number.forEach((el, index) => {
      this.field_pawn_number[index] = -1;
      this.isAbleToMove[index] = false;
    });
    if (this.state.ended) {
      return;
    }
    this.state.positions.forEach((player, index) => {
      player.forEach((pawn, pawnPlayerNumber) => {
        const position = board_positions[pawn];
        if (index == this.id) {
          this.field_pawn_number[pawn] = pawnPlayerNumber;
          if (this.playerStatus === this.CHOOSING_PAWN) {
            if (
              (pawn <= 15 &&
                (this.state.roll === 1 || this.state.roll === 6)) ||
              pawn >= 15
            ) {
              this.isAbleToMove[pawn] = true;
            }
          }
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

.possible-move {
  background-color: royalblue;
  border-radius: 100%;
  animation: blink 1s linear infinite;
}

.endscreen {
  z-index: 20;
  position: absolute;
  width: 100%;
  height: 100%;
  background-color: rgba(0, 0, 0, 0.5);
  display: flex;
  align-items: center;
  justify-content: center;
}

.centered-box {
  background-color: rgba(255, 255, 255, 1);
  font-size: 72px;
  width: 50rem;
  height: 20rem;
  font-family: "Segoe UI", Tahoma, Geneva, Verdana, sans-serif;
  display: flex;
  align-items: center;
  justify-content: center;
}

@keyframes blink {
  0% {
    opacity: 0;
  }
  100% {
    opacity: 0.5;
  }
}
.pinkHover {
  background-color: red;
  opacity: 0.5;
  border-radius: 100%;
}
.player-image {
  border-radius: 100%;
  width: 120px;
  height: 120px;
}
.player-container {
  margin-bottom: 2rem;
}
.dice {
  width: 120px;
  height: 120px;
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
