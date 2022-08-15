import React, { useEffect, useState } from 'react'
import Delete from '@mui/icons-material/Delete';
import CircularProgress from '@mui/material/CircularProgress';
import Box from '@mui/material/Box';
import WarehouseIcon from '@mui/icons-material/Warehouse';
import PersonIcon from '@mui/icons-material/Person';
import ItemUser from '../Item/ItemUser';
import Pagination from '../Pagination/Pagination';
import { useSelector } from 'react-redux';


function UserList() {
    const [pagination, setPagination] = useState({
        offset: 0,
        numberPerPage: 20,
        pageCount: 0,
    })

    const [userList, setUserList] = useState()

    const isSearchingSelector = useSelector((state) => state.user.isSearching)
    const userSelector = useSelector((state) => state.user.user)


    useEffect(() => {
        const fetchUser = async () => {
            let token = localStorage.getItem('token')
            try {
                const response = await fetch(`https://scm-tool.thanhpp.ninja/user`, {
                    headers: {
                        'Authorization': `Bearer ${token}`
                    }
                })
                if (!response) {
                    throw new Error(response.error.code);
                    return;
                }
                const data = await response.json()
                console.log(data)
                console.log(`Bearer ${token}`)

                setPagination(prev => ({
                    ...prev,
                    pageCount: !isSearchingSelector ? Math.ceil(data.data.length / prev.numberPerPage) : Math.ceil(userSelector.length / prev.numberPerPage)
                }))

                const userDisplay = !isSearchingSelector ? data.data.slice(pagination.offset, parseInt(pagination.offset) + parseInt(pagination.numberPerPage)) : userSelector.slice(pagination.offset, parseInt(pagination.offset) + parseInt(pagination.numberPerPage))

                setUserList(userDisplay)


            } catch (err) {
                alert(err)
            }
        }

        fetchUser()
    }, [pagination.offset, pagination.numberPerPage, pagination.pageCount, isSearchingSelector, userSelector])


    const numberPerPageHandle = (value) => {
        setPagination(prev => ({
            ...prev,
            numberPerPage: value
        }))
    }

    const handleClick = (e) => {
        const selected = e.selected
        const offset = selected * pagination.numberPerPage
        setPagination({ ...pagination, offset })
    }


    return (
        //bg-[#323259]
        <div className='bg-white rounded-md'>
            <div className='grid grid-cols-12 gap-3 text-black'>
                <div className='flex items-center col-span-1 my-[16px]'>
                    {/* <input type='checkbox' className='mr-[10px] leading-[24px] ml-[20px]' /> */}
                    <span className='mr-[10px] leading-[24px] ml-[20px]' >id</span>
                </div>
                <div className=' col-span-3 my-auto   truncate'>name</div>
                <div className=' col-span-5 my-auto truncate'>username</div>
                {/* <div className=' col-span-2 my-auto truncate'>location</div>
                <div className=' col-span-1 my-auto pl-[5px] flex justify-between truncate'>
                    <span className=' mr-[10px] rounded-sm'>Delete</span>
                    <span className=' mr-[10px] rounded-sm'>Edit</span>
                </div> */}
            </div>
            {userList && userList.map(user => <ItemUser id={user.id} key={user.id} name={user.name} username={user.username} />)}
            {/* <CircularProgress /> */}
            {/* </Box>} */}
            <Pagination getNumberPerPage={numberPerPageHandle} pageCount={pagination.pageCount} onPageChange={handleClick} />
        </div >
    )
}

export default UserList