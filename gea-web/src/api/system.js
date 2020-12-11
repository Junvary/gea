import service from "@/utils/request";

export const getSystemConfig = () => {
    return service({
        url: "/system/getSystemConfig",
        method: "post",
    });
};

export const setSystemConfig = (data) => {
    return service({
        url: "/system/setSystemConfig",
        method: "post",
        data,
    });
};

export const getSystemState = () => {
    return service({
        url: "/system/getServerInfo",
        method: "post",
        donNotShowLoading: true
    });
};