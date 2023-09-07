export interface RoomInfo {
    Platform: string;
    RoomID: string;
    Uname?: string;
    AreaName?: string;
    Title?: string;
    LiveStatus: number;
    LiveStartTime?: string;
    LiveStartTime2?: number;
    LiveTime?: string;
    State: number;
    StartTime?: string;
    EndTime?: string;
    RecordMode: boolean;
    RecordStartTime?: number;
    RecordEndTime?: number;
    RecordTime?: string;
    DecodeStartTime?: number;
    DecodeEndTime?: number;
    DecodeTime?: string;
    UploadStartTime?: number;
    UploadEndTime?: number;
    UploadTime?: string;
    DiskUsage?: number;
    NeedM4a?: boolean;
    Mp4Compress?: boolean;
    DivideByTitle?: boolean;
    CleanUpRegular?: boolean;
    SaveDuration?: string;
    AreaLock?: boolean;
    AreaLimit?: string;
    AutoRecord?: boolean;
    AutoUpload?: boolean;
    SaveDanmu?: boolean;
    Name?: string;
    Cookies?: string;
    port?: string;
  }
  
  export interface DeviceInfo {
    TotalDownload: string;
    FileNum: number;
  }
  
  export interface MonitorRoom {
    RoomID: string;
    UID: string;
    Uname: string;
    Title: string;
    Popularity: number;
    ParentID: string;
    ParentName: string;
    AreaID: string;
    AreaName: string;
    UserCover: string;
    LiveCover: string;
  }
  
  export interface submitInfo {
    platform: string;
    roomID: string;
    recordMode: boolean;
    startTime: string;
    endTime: string;
    autoRecord: boolean;
    autoUpload: boolean;
    needM4a: boolean;
    mp4Compress: boolean;
    divideByTitle: boolean;
    cleanUpRegular: boolean;
    saveDuration: string;
    areaLock: boolean;
    areaLimit: string;
    saveDanmu: boolean;
    name: string;
    cookies: string;
  }
  
  export interface liveStream {
    uname: string,
    roomID: string,
    streamURL: string,
    title: string
  }