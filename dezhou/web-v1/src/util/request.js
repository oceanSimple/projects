import axios from "axios";

// 创建axios实例
const requests = axios.create({
        baseURL: "/local",
        timeout: 5000
    }
);

// 请求拦截器
// requests.interceptors.request.use(
//     config => {
//         // 配置jwt token
//         const token = localStorage.getItem("jwtToken");
//         if (token) {
//             config.headers.Authorization = "Bearer " + token;
//         } else {
//             // 重定向到登录页面
//             window.location.href = "/#/login";
//             // 中止请求
//             return Promise.reject("require jwt token");
//         }
//         return config;
//     },
//     error => {
//         // 对请求错误做些什么
//         return Promise.reject(error);
//     }
// );

// 响应拦截器
requests.interceptors.response.use(
    response => {
        // 2xx开头的状态码都会进入这里
        // 这里只返回后端传递的数据data
        return response.data;
    },
    error => {
        // 对响应错误做点什么
        return Promise.reject(error);
    }
);

export default requests;