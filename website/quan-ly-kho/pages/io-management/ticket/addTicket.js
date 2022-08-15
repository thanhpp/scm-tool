import React, { useEffect, useState } from 'react'
import ArrowBackIcon from '@mui/icons-material/ArrowBack';
import { useRouter } from 'next/router';
import Loading from '../../../components/UI/Loading'
import axios from 'axios';
import { useSelector } from 'react-redux';

function AddTicket() {
    const [supplierId, setSupplierId] = useState()
    const [storageId, setStorageId] = useState()
    // const [status, setStatus] = useState()
    const [fee, setFee] = useState()
    const [billImage, setBillImage] = useState()
    const [productImage, setProductImage] = useState()
    const [receiveTime, setReceiveTime] = useState()

    const [itemSku, setItemSku] = useState()
    const [quantity, setQuantity] = useState()
    const [price, setPrice] = useState()

    // const [detail,setDetail] = useState()
    const [loading, setLoading] = useState(false)
    const router = useRouter()

    const [auth, setAuth] = useState()
    const [numberAddItem, setNumberAddItem] = useState([])


    useEffect(() => {
        if (!localStorage.getItem('token')) {
            router.push('/login')
            return;
        }
        const token = localStorage.getItem('token')
        setAuth(token)
    }, [])


    const supplierIdHandle = (e) => {
        setSupplierId(e.target.value)
    }

    const storageIdHandle = (e) => {
        setStorageId(e.target.value)
    }

    // const statusHandle = (e) => {
    //     setStatus(e.target.value)

    // }

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

    const itemSkuHandle = (e) => {
        setItemSku(e.target.value)
    }

    const quantityHandle = (e) => {
        setQuantity(e.target.value)
    }

    const priceHandle = (e) => {
        setPrice(e.target.value)
    }




    const addTicketHandle = async () => {
        setLoading(true)

        // const details = JSON.stringify([{
        //     item_sku: itemSku,
        //     buy_quantity: parseInt(quantity),
        //     buy_price: parseFloat(price)
        // }])
        const details = JSON.stringify(numberAddItem)

        // const escape = (str) => {
        //     return str
        //         .replace(/[\\]/g, "\\\\")
        //         .replace(/[\"]/g, '\\"')
        //         .replace(/[\/]/g, "\\/")
        //         .replace(/[\b]/g, "\\b")
        //         .replace(/[\f]/g, "\\f")
        //         .replace(/[\n]/g, "\\n")
        //         .replace(/[\r]/g, "\\r")
        //         .replace(/[\t]/g, "\\t");
        // };

        // const details = escape(x)
        // console.log(x)

        console.log(details)

        const formData = new FormData()

        formData.append('from_supplier_id', supplierId)
        formData.append('to_storage_id', storageId)
        // formData.append('status', status)
        formData.append('fee', fee)
        formData.append('bill_image_paths', billImage)
        formData.append('product_image_paths', productImage)
        formData.append('receive_time', receiveTime)
        formData.append('details', details)

        let token = localStorage.getItem('token')
        try {
            const res = await fetch('https://scm-tool.thanhpp.ninja/import_ticket', {
                method: 'POST',
                body: formData,
                headers: {
                    'Authorization': `Bearer ${token}`
                }
            })
            if (!res.ok) {
                throw new Error('something wrong');
                return;
            }

            const data = await res.json()
            if (data.error.code == 200) {
                alert(data.error.message)
                router.push('/io-management/ticket')
            }
            setLoading(false)
            console.log(data)
        } catch (err) {
            alert(err)
            setLoading(false)
        }
    }

    const addItemHandle = () => {
        setNumberAddItem([...numberAddItem, {
            item_sku: itemSku,
            buy_quantity: parseInt(quantity),
            buy_price: parseFloat(price)
        }])
        setItemSku('')
        setPrice('')
        setQuantity('')
        console.log(numberAddItem)
    }

    return (!auth ? <div></div> :
        <div className='bg-blue-500 w-full h-full min-h-screen   p-[36px]'>
            <div>
                <div onClick={() => router.back()} className=' bg-white rounded-[1000px] w-[30px] h-[30px] mb-[10px] cursor-pointer'>
                    <ArrowBackIcon className='text-black pl-[4px]' />
                </div>
                <div className=' text-[28px]'>
                    Add new ticket
                </div>
            </div>
            <div className='relative text-[#A16EFF] border-b border-white'>
                <span className='ml-[10px] cursor-pointer'>Info</span>
                <div className=' absolute w-[50px] h-[3px] bg-[#A16EFF]'></div>
            </div>
            <div className='my-[32px]'>
                <div>
                    {/* <div className='flex mb-[26px]'>
                        <div className='flex justify-end w-[175px] leading-[38px] mr-[20px]'>
                            <label htmlFor='id'>Id</label>
                        </div>
                        <input id='id' type='text' readOnly className='text-black h-[38px] w-[380px] pl-[15px] pr-[40px]' />
                    </div> */}
                    <div className='flex mb-[26px]'>
                        <div className='flex justify-end w-[175px] leading-[38px] mr-[20px]'>
                            <label htmlFor='supplierId'>supplier id</label>
                        </div>
                        <input id='supplierId' onChange={supplierIdHandle} type='text' className='text-black h-[38px] w-[380px] pl-[15px] pr-[40px]' />
                    </div>
                    <div className='flex mb-[26px]'>
                        <div className='flex justify-end w-[175px] leading-[38px] mr-[20px]'>
                            <label htmlFor='storageId'>storage id</label>
                        </div>
                        <input id='storageId' onChange={storageIdHandle} type='text' className='text-black h-[38px] w-[380px] pl-[15px] pr-[40px] truncate' />
                    </div>
                    {/* <div className='flex mb-[26px]'>
                        <div className='flex justify-end w-[175px] leading-[38px] mr-[20px]'>
                            <label htmlFor='status'>status</label>
                        </div>
                        <input id='status' onChange={statusHandle} type='text' className='text-black h-[38px] w-[380px] pl-[15px] pr-[40px] truncate' />
                    </div> */}
                    <div className='flex mb-[26px]'>
                        <div className='flex justify-end w-[175px] leading-[38px] mr-[20px]'>
                            <label htmlFor='fee'>fee</label>
                        </div>
                        <input id='fee' onChange={feeHandle} type='text' className='text-black h-[38px] w-[380px] pl-[15px] pr-[40px] truncate' />
                    </div>
                    <div className='flex mb-[26px]'>
                        <div className='flex justify-end w-[175px] leading-[38px] mr-[20px]'>
                            <label htmlFor='billImage'>bill image</label>
                        </div>
                        <input id='billImage' onChange={billImageHandle} type='file' className='text-black h-[38px] w-[380px] pl-[15px] pr-[40px] truncate' />
                    </div>
                    <div className='flex mb-[26px]'>
                        <div className='flex justify-end w-[175px] leading-[38px] mr-[20px]'>
                            <label htmlFor='productImage'>product image</label>
                        </div>
                        <input id='productImage' onChange={productImageHandle} type='file' className='text-black h-[38px] w-[380px] pl-[15px] pr-[40px] truncate' />
                    </div>
                    <div className='flex mb-[26px]'>
                        <div className='flex justify-end w-[175px] leading-[38px] mr-[20px]'>
                            <label htmlFor='receiveTime'>receive time</label>
                        </div>
                        <input id='receiveTime' onChange={receiveTimeHandle} type='text' className='text-black h-[38px] w-[380px] pl-[15px] pr-[40px] truncate' />
                    </div>
                    <div className='flex mb-[26px]'>
                        <div className='flex justify-end w-[175px] leading-[38px] mr-[20px]'>
                            <label htmlFor='itemSku'>item sku</label>
                        </div>
                        <input id='itemSku' onChange={itemSkuHandle} value={itemSku || ''} type='text' className='text-black h-[38px] w-[380px] pl-[15px] pr-[40px] truncate' />
                    </div>
                    <div className='flex mb-[26px]'>
                        <div className='flex justify-end w-[175px] leading-[38px] mr-[20px]'>
                            <label htmlFor='quantity'>quantity</label>
                        </div>
                        <input id='quantity' onChange={quantityHandle} value={quantity || ''} type='text' className='text-black h-[38px] w-[380px] pl-[15px] pr-[40px] truncate' />
                    </div>
                    <div className='flex mb-[26px]'>
                        <div className='flex justify-end w-[175px] leading-[38px] mr-[20px]'>
                            <label htmlFor='price'>price</label>
                        </div>
                        <input id='price' onChange={priceHandle} value={price || ''} type='text' className='text-black h-[38px] w-[380px] pl-[15px] pr-[40px] truncate' />
                    </div>
                    <div className='flex mb-[26px]'>
                        <div onClick={addItemHandle} className='flex justify-end w-[175px] leading-[38px] ml-[87px] cursor-pointer'>
                            <div className=' border-white border border-collapse rounded-md'>Add item</div>
                        </div>
                    </div>
                    {numberAddItem && numberAddItem.map((item) => (
                        <div key={item.item_sku} className='flex mb-[26px]'>
                            <div onClick={addItemHandle} className='flex justify-end  leading-[38px] ml-[195px] cursor-pointer'>
                                <div className='border-white border border-collapse p-2 '>{'sku: ' + item.item_sku}</div>
                                <div className='border-white border border-collapse p-2 '>{'quantity: ' + item.buy_quantity}</div>
                                <div className='border-white border border-collapse p-2 '>{'price: ' + item.buy_price}</div>
                            </div>
                        </div>
                    ))}
                    {/* <div className='flex mb-[26px]'>
                        <div className='flex justify-end w-[175px] leading-[38px] mr-[20px]'>
                            <label htmlFor='sendTime'>send time</label>
                        </div>
                        <input id='sendTime' onChange={sendTimeHandle} type='text' className='text-black h-[38px] w-[380px] pl-[15px] pr-[40px] truncate' />
                    </div> */}
                    <div className='flex mb-[26px]'>
                        <div onClick={addTicketHandle} className='h-[38px] text-center hover:opacity-80 rounded-md ml-[200px] px-[12px] py-[6px] bg-transparent border border-white cursor-pointer'>
                            <button className='text-white ' >Add new ticket</button>
                        </div>
                    </div>
                </div>
            </div>
            {loading ? <Loading /> : ''}
        </div>
    )
}

export default AddTicket