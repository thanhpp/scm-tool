import React, { useState } from 'react'
import PowerSettingsNew from '@mui/icons-material/PowerSettingsNew';
import Delete from '@mui/icons-material/Delete';
import InboxIcon from '@mui/icons-material/Inbox';
import Link from 'next/link';
import EditIcon from '@mui/icons-material/Edit';
import Modal from '../UI/Modal';
import Loading from '../UI/Loading';
import { useRouter } from 'next/router';

function ItemProduct({ sku, name, desc, itemTypeId, images }) {
    // const [skuUpdate, setSkuUpdate] = useState(sku)
    const [nameUpdate, setNameUpdate] = useState(name)
    const [descUpdate, setDescUpdate] = useState(desc)
    const [itemTypeIdUpdate, setItemTypeIdUpdate] = useState(itemTypeId)
    const [imageUpdate, setImageUpdate] = useState(images)
    const [loading, setLoading] = useState(false)

    const router = useRouter()

    const [showModal, setShowModal] = useState(false)

    // const skuHandle = (e) => {
    //     setSkuUpdate(e.target.value)
    // }

    const nameHandle = (e) => {
        setNameUpdate(e.target.value)
    }

    const descHandle = (e) => {
        setDescUpdate(e.target.value)
    }

    const itemTypeIdHandle = (e) => {
        setItemTypeIdUpdate(e.target.value)
    }

    const imageHandle = (e) => {
        setImageUpdate(e.target.files[0])
    }


    const updateHandle = async () => {
        setLoading(true)

        const domainImage = 'https://scm-tool.thanhpp.ninja/files/'
        const domainImageLength = domainImage.length
        const imageDelete = images[0].slice(domainImageLength)

        const dataForm = new FormData()

        // dataForm.append('sku', skuUpdate)
        dataForm.append('name', nameUpdate || name)
        dataForm.append('desc', descUpdate || desc)
        dataForm.append('item_type_id', itemTypeIdUpdate || itemTypeId)
        dataForm.append('new_images', imageUpdate || images)
        if (imageDelete) {
            dataForm.append('delete_images', imageDelete)
        }
        let token = localStorage.getItem('token')
        try {
            const res = await fetch(` https://scm-tool.thanhpp.ninja/item/${sku}`, {
                method: 'PUT',
                body: dataForm,
                headers: {
                    'Content-type': 'application/json',
                    "Authorization": `Bearer ${token}`
                }
            })

            if (!res.ok) {
                throw new Error('edit failed')
                return;
            }

            const data = await res.json()
            console.log(images)
            router.reload(window.location.pathname)
            setLoading(false)
        } catch (err) {
            alert(err)
            setLoading(false)
        }
    }

    return (
        <div className=' grid grid-cols-12 gap-3 text-black  border-black border-t'>
            <div className='flex items-center col-span-1 my-[16px]'>
                {/* <input type='checkbox' className='mr-[10px] leading-[24px] ml-[20px]' /> */}
                {/* <InboxIcon className='mr-[10px] leading-[24px] ml-[20px]' /> */}
                <span className='mr-[10px] leading-[24px] ml-[20px]' >{sku}</span>
            </div>
            <Link href={`/products-management/${sku}`}><div className=' col-span-3 my-auto hover:border-b border-black cursor-pointer  truncate'>{name}</div></Link>
            <div className=' col-span-5 my-auto truncate'>{desc}</div>
            <div className=' col-span-2 my-auto truncate'>{itemTypeId}</div>
            <div className=' col-span-1 my-auto pl-[5px] flex justify-between truncate'>
                {/* <span className=' mr-[10px] rounded-sm cursor-pointer hover:opacity-80'><Delete className='' /></span> */}
                <span onClick={() => setShowModal(true)} className=' mr-[10px] rounded-sm cursor-pointer hover:opacity-80'><EditIcon className='' /></span>
            </div>
            <Modal show={showModal} onClose={() => setShowModal(false)}>
                <div>
                    <div className='flex mb-[26px] '>
                        <div className='flex justify-end w-[175px] leading-[38px] mr-[20px] '>
                            <label htmlFor='sku'>Sku<span className='text-[#d13143]'>*</span></label>
                        </div>
                        <input id='sku' type='text' onChange={() => { }} value={sku} className='text-black h-[38px] w-[380px] pl-[15px] pr-[40px]' />
                    </div>
                    <div className='flex mb-[26px]'>
                        <div className='flex justify-end w-[175px] leading-[38px] mr-[20px]'>
                            <label htmlFor='name'>Name<span className='text-[#d13143]'>*</span></label>
                        </div>
                        <input id='name' onChange={nameHandle} type='text' value={nameUpdate} className='text-black h-[38px] w-[380px] pl-[15px] pr-[40px]' />
                    </div>
                    <div className='flex mb-[26px]'>
                        <div className='flex justify-end w-[175px] leading-[38px] mr-[20px]'>
                            <label htmlFor='description'>Description<span className='text-[#d13143]'>*</span></label>
                        </div>
                        <input id='description' onChange={descHandle} value={descUpdate} type='text' className='text-black h-[38px] w-[380px] pl-[15px] pr-[40px] truncate' />
                    </div>
                    <div className='flex mb-[26px]'>
                        <div className='flex justify-end w-[175px] leading-[38px] mr-[20px]'>
                            <label htmlFor='itemTypeId'>Item type id<span className='text-[#d13143]'>*</span></label>
                        </div>
                        <input id='itemTypeId' type='text' value={itemTypeIdUpdate} onChange={itemTypeIdHandle} className='text-black h-[38px] w-[380px] pl-[15px] pr-[40px] truncate' />
                    </div>
                    <div className='flex mb-[26px]'>
                        <div className=' ml-[125px]'><a href={images} target="_blank" rel="noreferrer">View image</a></div>
                    </div>
                    <div className='flex mb-[26px]'>
                        <div className='flex justify-end w-[175px] leading-[38px] mr-[20px]'>
                            <label htmlFor='image'>Image<span className='text-[#d13143]'>*</span></label>
                        </div>
                        <input id='image' type='file' onChange={imageHandle} className='text-black h-[38px] w-[380px] pl-[15px] pr-[40px] truncate' />
                    </div>
                </div>
                <div className='flex '>
                    <div onClick={updateHandle} className='h-[38px] text-center hover:opacity-80 rounded-md ml-[200px] px-[12px] py-[6px]  border border-white bg-blue-500 cursor-pointer'>
                        <button className='text-black ' >Update Product</button>
                    </div>
                </div>
                {loading && <Loading />}
            </Modal>
        </div>
    )
}

export default ItemProduct