import service from '@/utils/request'

export const getPayCode = (data) => {
    return service({
        url: '/wxpay/getPayCode',
        method: 'post',
        data
    })
}

export const getOrderById = (params) => {
    return service({
        url: '/wxpay/getOrderById',
        method: 'get',
        params,
        doNotShowLoading: true
    })
}


