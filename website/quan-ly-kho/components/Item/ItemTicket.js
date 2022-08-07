import React, { useState } from 'react'
import PowerSettingsNew from '@mui/icons-material/PowerSettingsNew';
import Delete from '@mui/icons-material/Delete';
import InboxIcon from '@mui/icons-material/Inbox';
import Link from 'next/link';
import EditIcon from '@mui/icons-material/Edit';
import Modal from '../UI/Modal';
import Loading from '../UI/Loading';
import { useRouter } from 'next/router';

function ItemTicket({ id, supplierId, storageId, status, fee, billImage, productImage, receiveTime }) {
    const [idUpdate, setIdUpdate] = useState(id)
    const [supplierIdUpdate, setSupplierIdUpdate] = useState(supplierId)
    const [storageIdUpdate, setStorageIdUpdate] = useState(storageId)
    const [statusUpdate, setStatusUpdate] = useState(status)
    const [feeUpdate, setFeeUpdate] = useState(fee)
    const [billImageUpdate, setBillImageUpdate] = useState(billImage)
    const [productImageUpdate, setProductImageUpdate] = useState(productImage)
    const [receiveTimeUpdate, setReceiveTimeUpdate] = useState(receiveTime)
    // const [sendTimeUpdate, setSendTimeUpdate] = useState(sendTime)
    const [loading, setLoading] = useState(false)

    const router = useRouter()

    const [showModal, setShowModal] = useState(false)

    // const idHandle = (e) => {
    //     setIdUpdate(e.target.value)
    // }

    const supplierIdHandle = (e) => {
        setSupplierIdUpdate(e.target.value)
    }

    const storageIdHandle = (e) => {
        setStorageIdUpdate(e.target.value)
    }

    // const statusHandle = (e) => {
    //     setStatusUpdate(e.target.value)
    // }

    const feeHandle = (e) => {
        setFeeUpdate(e.target.value)
    }

    const billImageHandle = (e) => {
        setBillImageUpdate(e.target.files[0])
    }

    const productImageHandle = (e) => {
        setProductImageUpdate(e.target.files[0])
    }

    const receiveTimeHandle = (e) => {
        setReceiveTimeUpdate(e.target.value)
    }

    // const sendTimeHandle = (e) => {
    //     setSendTimeUpdate(e.target.value)
    // }


    const updateHandle = async () => {
        setLoading(true)


        try {
            const res = await fetch(` https://scm-tool.thanhpp.ninja/import_ticket/${id}`, {
                method: 'PUT',
                body: JSON.stringify({
                    name: nameUpdate,
                    desc: descUpdate
                }),
                headers: {
                    'Content-type': 'application/json'
                }
            })

            if (!res) {
                throw new Error('edit failed')
                return;
            }

            const data = await res.json()
            if (data.error.code == 200) {
                console.log('edit success')
                router.reload(window.location.pathname)
            }
            setLoading(false)
        } catch (err) {
            console.log('edit fail')
            setLoading(false)
        }
    }

    return (
        <div className=' grid grid-cols-12 gap-3 text-black  border-black border-t'>
            <div className='flex items-center col-span-1 my-[16px]'>
                <span className='mr-[10px] leading-[24px] ml-[20px]' >{id}</span>
            </div>
            <div className='flex items-center col-span-1 my-[16px]'>
                <span className='mr-[10px] leading-[24px] ml-[20px]' >{supplierId}</span>
            </div>
            <div className='flex items-center col-span-1 my-[16px]'>
                <span className='mr-[10px] leading-[24px] ml-[20px]' >{storageId}</span>
            </div>
            <Link href={`/io-management/ticket/${id}`}><div className=' col-span-6 my-auto hover:border-b border-black cursor-pointer  truncate'>{status}</div></Link>
            <div className=' col-span-2 my-auto truncate'>{fee}</div>
            <div className=' col-span-1 my-auto pl-[5px] flex justify-between truncate'>
                <span className=' mr-[10px] rounded-sm cursor-pointer hover:opacity-80'><Delete className='' /></span>
                <span onClick={() => setShowModal(true)} className=' mr-[10px] rounded-sm cursor-pointer hover:opacity-80'><EditIcon className='' /></span>
            </div>
            <Modal show={showModal} onClose={() => setShowModal(false)}>
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
                        <input id='supplierId' onChange={supplierIdHandle} type='text' value={supplierIdUpdate} className='text-black h-[38px] w-[380px] pl-[15px] pr-[40px]' />
                    </div>
                    <div className='flex mb-[26px]'>
                        <div className='flex justify-end w-[175px] leading-[38px] mr-[20px]'>
                            <label htmlFor='storageId'>storage id</label>
                        </div>
                        <input id='storageId' onChange={storageIdHandle} value={storageIdUpdate} type='text' className='text-black h-[38px] w-[380px] pl-[15px] pr-[40px] truncate' />
                    </div>
                    <div className='flex mb-[26px]'>
                        <div className='flex justify-end w-[175px] leading-[38px] mr-[20px]'>
                            <label htmlFor='status'>status</label>
                        </div>
                        <input id='status' readOnly value={statusUpdate} type='text' className='text-black h-[38px] w-[380px] pl-[15px] pr-[40px] truncate' />
                    </div>
                    <div className='flex mb-[26px]'>
                        <div className='flex justify-end w-[175px] leading-[38px] mr-[20px]'>
                            <label htmlFor='fee'>fee</label>
                        </div>
                        <input id='fee' onChange={feeHandle} value={feeUpdate} type='text' className='text-black h-[38px] w-[380px] pl-[15px] pr-[40px] truncate' />
                    </div>
                    <div className='flex mb-[26px]'>
                        <div className='flex justify-end w-[175px] leading-[38px] mr-[20px]'>
                            <label htmlFor='billImage'>bill image</label>
                        </div>
                        <input id='billImage' onChange={billImageHandle} value={billImageUpdate} type='text' className='text-black h-[38px] w-[380px] pl-[15px] pr-[40px] truncate' />
                    </div>
                    <div className='flex mb-[26px]'>
                        <div className='flex justify-end w-[175px] leading-[38px] mr-[20px]'>
                            <label htmlFor='productImage'>product image</label>
                        </div>
                        <input id='productImage' onChange={productImageHandle} value={productImageUpdate} type='text' className='text-black h-[38px] w-[380px] pl-[15px] pr-[40px] truncate' />
                    </div>
                    <div className='flex mb-[26px]'>
                        <div className='flex justify-end w-[175px] leading-[38px] mr-[20px]'>
                            <label htmlFor='receiveTime'>receive time</label>
                        </div>
                        <input id='receiveTime' onChange={receiveTimeHandle} value={receiveTimeUpdate} type='text' className='text-black h-[38px] w-[380px] pl-[15px] pr-[40px] truncate' />
                    </div>
                    {/* <div className='flex mb-[26px]'>
                        <div className='flex justify-end w-[175px] leading-[38px] mr-[20px]'>
                            <label htmlFor='sendTime'>send time</label>
                        </div>
                        <input id='sendTime' onChange={sendTimeHandle} value={sendTimeUpdate} type='text' className='text-black h-[38px] w-[380px] pl-[15px] pr-[40px] truncate' />
                    </div> */}
                </div>
                <div className='flex '>
                    <div onClick={updateHandle} className='h-[38px] text-center hover:opacity-80 rounded-md ml-[200px] px-[12px] py-[6px]  border border-white bg-blue-500 cursor-pointer'>
                        <button className='text-black ' >Update Ticket</button>
                    </div>
                </div>
                {loading && <Loading />}
            </Modal>
        </div>
    )
}

export default ItemTicket