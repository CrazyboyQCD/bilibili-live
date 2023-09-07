function state2Type(state: number): string {
  if (state == 1 || state == 2) {
    return "";
  } else if (state == 3 || state == 5 || state == 8) {
    return "success";
  } else if (state == 4 || state == 7) {
    return "warning";
  } else {
    return "danger";
  }
}

function state2Name(state: number): string {
  switch (state) {
    case 1:
      return "正在监听";
    case 2:
      return "等待重连";
    case 3:
      return "录制中";
    case 4:
      return "等待转码";
    case 5:
      return "转码中";
    case 6:
      return "转码结束";
    case 7:
      return "等待上传";
    case 8:
      return "上传中";
    case 9:
      return "上传结束";
    case 10:
      return "停止监听";
  }
  return "未知";
}

function status2Type(status: number): string {
  switch (status) {
    case 0:
      return "default";
    case 1:
      return "success";
    case 2:
      return "warning";
  }
  return "error";
}

function status2Name(status: number): string {
  switch (status) {
    case 0:
      return "未开播";
    case 1:
      return "直播中";
    case 2:
      return "录播中";
  }
  return "封禁中";
}

export { state2Type, state2Name, status2Type, status2Name };
