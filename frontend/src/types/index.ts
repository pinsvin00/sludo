export class LudoGM {
  canvas: HTMLCanvasElement = new HTMLCanvasElement();
  canvasCtx: CanvasRenderingContext2D = new CanvasRenderingContext2D();
  ludoImage?: CanvasImageData;
}

export enum playerStatus {
  NOT_READY,
  READY,
  ROLLING_DICE,
  CHOOSING_PAWN,
  AFK,
}

export class GameState {
  id: number;
  positions: Array<Array<number>> = [];
  started: boolean;
  names: Array<string> = [];
  players_status: Array<number> = [];
  player_count: number;
  turn = -2;
  roll: number;
}

//DO 15 pozycji pola startowe
//DO 45 index pola gry
//DO KONCA to bedzie reszta pol
export const board_positions = [
  {
    x: 53,
    y: 50,
  },
  {
    x: 121,
    y: 49,
  },
  {
    x: 50,
    y: 121,
  },
  {
    x: 119,
    y: 120,
  },
  {
    x: 674,
    y: 46,
  },
  {
    x: 749,
    y: 50,
  },
  {
    x: 680,
    y: 119,
  },
  {
    x: 749,
    y: 124,
  },
  {
    x: 680,
    y: 680,
  },
  {
    x: 750,
    y: 676,
  },
  {
    x: 676,
    y: 751,
  },
  {
    x: 749,
    y: 746,
  },
  {
    x: 50,
    y: 676,
  },
  {
    x: 118,
    y: 675,
  },
  {
    x: 45,
    y: 747,
  },
  {
    x: 119,
    y: 743,
  },
  {
    x: 51,
    y: 328,
  },
  {
    x: 119,
    y: 328,
  },
  {
    x: 193,
    y: 328,
  },
  {
    x: 258,
    y: 326,
  },
  {
    x: 327,
    y: 329,
  },
  {
    x: 329,
    y: 261,
  },
  {
    x: 332,
    y: 191,
  },
  {
    x: 332,
    y: 120,
  },
  {
    x: 330,
    y: 49,
  },
  {
    x: 399,
    y: 48,
  },
  {
    x: 472,
    y: 50,
  },
  {
    x: 470,
    y: 119,
  },
  {
    x: 468,
    y: 188,
  },
  {
    x: 472,
    y: 258,
  },
  {
    x: 471,
    y: 325,
  },
  {
    x: 538,
    y: 328,
  },
  {
    x: 604,
    y: 327,
  },
  {
    x: 679,
    y: 329,
  },
  {
    x: 746,
    y: 328,
  },
  {
    x: 748,
    y: 397,
  },
  {
    x: 747,
    y: 468,
  },
  {
    x: 677,
    y: 468,
  },
  {
    x: 608,
    y: 469,
  },
  {
    x: 541,
    y: 467,
  },
  {
    x: 466,
    y: 466,
  },
  {
    x: 467,
    y: 539,
  },
  {
    x: 467,
    y: 609,
  },
  {
    x: 469,
    y: 675,
  },
  {
    x: 470,
    y: 748,
  },
  {
    x: 396,
    y: 749,
  },
  {
    x: 330,
    y: 747,
  },
  {
    x: 333,
    y: 674,
  },
  {
    x: 329,
    y: 606,
  },
  {
    x: 326,
    y: 537,
  },
  {
    x: 328,
    y: 468,
  },
  {
    x: 258,
    y: 473,
  },
  {
    x: 190,
    y: 469,
  },
  {
    x: 120,
    y: 466,
  },
  {
    x: 50,
    y: 467,
  },
  {
    x: 40,
    y: 390,
  },
  {
    x: 122,
    y: 399,
  },
  {
    x: 190,
    y: 397,
  },
  {
    x: 255,
    y: 398,
  },
  {
    x: 327,
    y: 399,
  },
  {
    x: 397,
    y: 123,
  },
  {
    x: 398,
    y: 189,
  },
  {
    x: 401,
    y: 259,
  },
  {
    x: 400,
    y: 322,
  },
  {
    x: 673,
    y: 393,
  },
  {
    x: 608,
    y: 396,
  },
  {
    x: 540,
    y: 390,
  },
  {
    x: 484,
    y: 391,
  },
  {
    x: 401,
    y: 666,
  },
  {
    x: 395,
    y: 607,
  },
  {
    x: 403,
    y: 545,
  },
  {
    x: 399,
    y: 473,
  },
];
//yellow, green, blue, red
export const player_colors = ["#850000", "#005eab", "#238500", "#e5ff00"];
