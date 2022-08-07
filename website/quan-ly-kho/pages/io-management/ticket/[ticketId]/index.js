import { useRouter } from 'next/router'
import React, { useEffect, useLayoutEffect, useState } from 'react'
import ArrowBackIcon from '@mui/icons-material/ArrowBack';
import Loading from '../../../../components/UI/Loading';

function TicketId() {

    const [ticketDetail, setTicketDetail] = useState()

    const [id, setId] = useState()
    const [supplierId, setSupplierId] = useState()
    const [storageId, setStorageId] = useState()
    const [status, setStatus] = useState()
    const [fee, setFee] = useState()
    const [billImage, setBillImage] = useState()
    const [productImage, setProductImage] = useState()
    const [receiveTime, setReceiveTime] = useState()

    const [billImageDisplay, setBillImageDisplay] = useState()
    const [productImageDisplay, setProductImageDisplay] = useState()
    const [loading, setLoading] = useState(false)
    const router = useRouter()

    const { ticketId } = router.query


    const fetchTicket = async () => {
        setLoading(true)

        try {
            const res = await fetch(`https://scm-tool.thanhpp.ninja/import_ticket/${ticketId}`)

            if (!res.ok) {
                throw new Error('can not fetch data')
                return;
            }

            const data = await res.json()

            const ticketData = data.data

            setId(ticketData.id)
            setSupplierId(ticketData.from_supplier_id)
            setStorageId(ticketData.to_storage_id)
            setStatus(ticketData.status)
            setFee(ticketData.fee)
            setBillImage(ticketData.bill_image_paths)
            setProductImage(ticketData.product_image_paths)
            setReceiveTime(ticketData.receive_time)

            setBillImageDisplay(ticketData.bill_image_paths && ticketData.bill_image_paths[0])
            setProductImageDisplay(ticketData.product_image_paths && ticketData.product_image_paths[0])


            console.log(data.data)

            setLoading(false)
        } catch (err) {
            console.log(err)
            setLoading(false)
        }

        setLoading(false)
    }

    useEffect(() => {
        if (!router.isReady) {
            return;
        }

        fetchTicket()
    }, [router.isReady])


    const supplierIdHandle = (e) => {
        setSupplierId(e.target.value)
    }

    const storageIdHandle = (e) => {
        setStorageId(e.target.value)
    }

    const statusHandle = (e) => {
        setStatus(e.target.value)
    }

    const feeHandle = (e) => {
        setFee(e.target.value)
    }

    const billImageHandle = (e) => {
        setBillImage(e.target.files[0])
    }

    const productImageHandle = (e) => {
        setProductImage(e.target.files[0])
    }

    const receiveTimeHandle = (e) => {
        setReceiveTime(e.target.value)
    }

    // const updateTicketHandle = async () => {
    //     setLoading(true)

    //     const formData = new FormData()

    //     formData.append('from_supplier_id', supplierId)
    //     formData.append('to_storage_id', storageId)
    //     formData.append('status', status)
    //     formData.append('fee', fee)
    //     formData.append('bill_image_paths', billImage)
    //     formData.append('product_image_paths', productImage)
    //     formData.append('receive_time', receiveTime)

    //     try {
    //         const res = await fetch(` https://scm-tool.thanhpp.ninja/import_ticket/${ticketId}`, {
    //             method: 'PUT',
    //             body: formData
    //             // headers: {
    //             //     "Content-type": "multipart/form-data"
    //             // }
    //         })
    //         if (!res.ok) {
    //             throw new Error('something wrong');
    //             return;
    //         }

    //         const data = await res.json()
    //         if (data.error.code == 200) {
    //             router.push('/io-management/ticket')
    //         }
    //         setLoading(false)
    //         console.log(data)
    //     } catch (err) {
    //         console.log(err)
    //         setLoading(false)
    //     }
    // }

    return (
        <div className='bg-blue-500 w-full h-full min-h-screen   p-[36px]'>
            <div>
                <div onClick={() => router.back()} className=' bg-white rounded-[1000px] w-[30px] h-[30px] mb-[10px] cursor-pointer'>
                    <ArrowBackIcon className='text-black pl-[4px]' />
                </div>
                <div className=' text-[28px]'>
                    {/* {name?.name} */}
                </div>
            </div>
            <div className='relative text-[#A16EFF] border-b border-white'>
                <span className='ml-[10px] cursor-pointer'>Info</span>
                <div className=' absolute w-[50px] h-[3px] bg-[#A16EFF]'></div>
            </div>
            <div className='my-[32px]'>
                <div>
                    <div className='flex mb-[26px]'>
                        <div className='flex justify-end w-[175px] leading-[38px] mr-[20px]'>
                            <label htmlFor='id'>Id</label>
                        </div>
                        <input id='id' type='text' defaultValue={id} readOnly className='text-black h-[38px] w-[380px] pl-[15px] pr-[40px]' />
                    </div>
                    <div className='flex mb-[26px]'>
                        <div className='flex justify-end w-[175px] leading-[38px] mr-[20px]'>
                            <label htmlFor='supplierId'>supplier id</label>
                        </div>
                        <input id='supplierId' onChange={supplierIdHandle} readOnly type='text' value={supplierId ?? ""} className='text-black h-[38px] w-[380px] pl-[15px] pr-[40px]' />
                    </div>
                    <div className='flex mb-[26px]'>
                        <div className='flex justify-end w-[175px] leading-[38px] mr-[20px]'>
                            <label htmlFor='storageId'>storage id</label>
                        </div>
                        <input id='storageId' onChange={storageIdHandle} readOnly value={storageId ?? ""} type='text' className='text-black h-[38px] w-[380px] pl-[15px] pr-[40px] truncate' />
                    </div>
                    <div className='flex mb-[26px]'>
                        <div className='flex justify-end w-[175px] leading-[38px] mr-[20px]'>
                            <label htmlFor='status'>status</label>
                        </div>
                        <input id='status' onChange={statusHandle} readOnly value={status ?? ""} type='text' className='text-black h-[38px] w-[380px] pl-[15px] pr-[40px] truncate' />
                    </div>
                    <div className='flex mb-[26px]'>
                        <div className='flex justify-end w-[175px] leading-[38px] mr-[20px]'>
                            <label htmlFor='fee'>fee</label>
                        </div>
                        <input id='fee' onChange={feeHandle} readOnly value={fee ?? ""} type='text' className='text-black h-[38px] w-[380px] pl-[15px] pr-[40px] truncate' />
                    </div>
                    <div className='flex mb-[26px]'>
                        <div className='flex justify-end w-[175px] leading-[38px] mr-[20px]'>
                            <label htmlFor='billImage'>bill image</label>
                        </div>
                        <div>
                            <div className=' max-h-full max-w-[500px]'><img src={billImageDisplay} /></div>
                            {/* <div><input id='billImage' onChange={billImageHandle} type='file' className='text-black h-[38px] w-[380px] pl-[15px] pr-[40px] truncate' /></div> */}
                        </div>
                    </div>
                    <div className='flex mb-[26px]'>
                        <div className='flex justify-end w-[175px] leading-[38px] mr-[20px]'>
                            <label htmlFor='productImage'>product image</label>
                        </div>
                        <div>
                            <div className=' max-h-full max-w-[500px]'><img src={productImageDisplay} /></div>
                            {/* <div><input id='productImage' onChange={productImageHandle} type='file' className='text-black h-[38px] w-[380px] pl-[15px] pr-[40px] truncate' /></div> */}
                        </div>
                    </div>
                    <div className='flex mb-[26px]'>
                        <div className='flex justify-end w-[175px] leading-[38px] mr-[20px]'>
                            <label htmlFor='receiveTime'>receive time</label>
                        </div>
                        <input id='receiveTime' onChange={receiveTimeHandle} value={receiveTime ?? ""} type='text' className='text-black h-[38px] w-[380px] pl-[15px] pr-[40px] truncate' />
                    </div>
                    {/* <div className='flex mb-[26px]'>
                        <div className='flex justify-end w-[175px] leading-[38px] mr-[20px]'>
                            <label htmlFor='sendTime'>send time</label>
                        </div>
                        <input id='sendTime' onChange={sendTimeHandle} value={sendTime} type='text' className='text-black h-[38px] w-[380px] pl-[15px] pr-[40px] truncate' />
                    </div> */}
                    {/* <div className='flex mb-[26px]'>
                        <div onClick={updateTicketHandle} className='h-[38px] text-center hover:opacity-80 rounded-md ml-[200px] px-[12px] py-[6px] bg-transparent border border-white cursor-pointer'>
                            <button className='text-white ' >Update ticket</button>
                        </div>
                    </div> */}
                </div>
            </div>
            {loading ? <Loading /> : ''}
        </div>
    )
}

export default TicketId