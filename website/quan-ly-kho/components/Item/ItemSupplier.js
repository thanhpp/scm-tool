import React, { useState } from 'react'
import PowerSettingsNew from '@mui/icons-material/PowerSettingsNew';
import Delete from '@mui/icons-material/Delete';
import InboxIcon from '@mui/icons-material/Inbox';
import Link from 'next/link';
import EditIcon from '@mui/icons-material/Edit';
import Modal from '../UI/Modal';
import Loading from '../UI/Loading';
import { useRouter } from 'next/router';

function ItemSupplier({ id, name, phone, email }) {
    const [idUpdate, setIdUpdate] = useState(id)
    const [nameUpdate, setNameUpdate] = useState(name)
    const [phoneUpdate, setPhoneUpdate] = useState(phone)
    const [emailUpdate, setEmailUpdate] = useState(email)
    const [loading, setLoading] = useState(false)

    const router = useRouter()

    const [showModal, setShowModal] = useState(false)

    const idHandle = (e) => {
        setIdUpdate(e.target.value)
    }

    const nameHandle = (e) => {
        setNameUpdate(e.target.value)
    }

    const emailHandle = (e) => {
        setEmailUpdate(e.target.value)
    }

    const phoneHandle = (e) => {
        setPhoneUpdate(e.target.value)
    }

    const updateHandle = async () => {
        setLoading(true)


        try {
            const res = await fetch(` https://scm-tool.thanhpp.ninja/supplier/${id}`, {
                method: 'PUT',
                body: JSON.stringify({
                    name: nameUpdate,
                    email: emailUpdate,
                    phone: phoneUpdate
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
                {/* <input type='checkbox' className='mr-[10px] leading-[24px] ml-[20px]' /> */}
                {/* <InboxIcon className='mr-[10px] leading-[24px] ml-[20px]' /> */}
                <span className='mr-[10px] leading-[24px] ml-[20px]' >{id}</span>
            </div>
            <div className=' col-span-4 my-auto  truncate'>{name}</div>
            {/* <Link href={`/products-management/supplier/${id}`}><div className=' col-span-4 my-auto hover:border-b border-black cursor-pointer  truncate'>{name}</div></Link> */}
            <div className=' col-span-4 my-auto truncate'>{email}</div>
            <div className=' col-span-2 my-auto truncate'>{phone}</div>
            <div className=' col-span-1 my-auto pl-[5px] flex justify-between truncate'>
                <span className=' mr-[10px] rounded-sm cursor-pointer hover:opacity-80'><Delete className='' /></span>
                <span onClick={() => setShowModal(true)} className=' mr-[10px] rounded-sm cursor-pointer hover:opacity-80'><EditIcon className='' /></span>
            </div>
            <Modal show={showModal} onClose={() => setShowModal(false)}>
                <div>
                    <div className='flex mb-[26px]'>
                        <div className='flex justify-end w-[175px] leading-[38px] mr-[20px]'>
                            <label htmlFor='id'>Id<span className='text-[#d13143]'>*</span></label>
                        </div>
                        <input id='id' onChange={idHandle} type='text' value={idUpdate} className='text-black h-[38px] w-[380px] pl-[15px] pr-[40px]' />
                    </div>
                    <div className='flex mb-[26px]'>
                        <div className='flex justify-end w-[175px] leading-[38px] mr-[20px]'>
                            <label htmlFor='name'>Name<span className='text-[#d13143]'>*</span></label>
                        </div>
                        <input id='name' onChange={nameHandle} type='text' value={nameUpdate} className='text-black h-[38px] w-[380px] pl-[15px] pr-[40px]' />
                    </div>
                    <div className='flex mb-[26px]'>
                        <div className='flex justify-end w-[175px] leading-[38px] mr-[20px]'>
                            <label htmlFor='email'>Email<span className='text-[#d13143]'>*</span></label>
                        </div>
                        <input id='email' onChange={emailHandle} value={emailUpdate} type='text' className='text-black h-[38px] w-[380px] pl-[15px] pr-[40px] truncate' />
                    </div>
                    <div className='flex mb-[26px]'>
                        <div className='flex justify-end w-[175px] leading-[38px] mr-[20px]'>
                            <label htmlFor='phone'>Phone<span className='text-[#d13143]'>*</span></label>
                        </div>
                        <input id='phone' onChange={phoneHandle} value={phoneUpdate} type='text' className='text-black h-[38px] w-[380px] pl-[15px] pr-[40px] truncate' />
                    </div>
                </div>
                <div className='flex '>
                    <div onClick={updateHandle} className='h-[38px] text-center hover:opacity-80 rounded-md ml-[200px] px-[12px] py-[6px]  border border-white bg-blue-500 cursor-pointer'>
                        <button className='text-black ' >Update supplier</button>
                    </div>
                </div>
                {loading && <Loading />}
            </Modal>
        </div>
    )
}

export default ItemSupplier