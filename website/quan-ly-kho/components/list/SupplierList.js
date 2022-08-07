import React, { useEffect, useState } from 'react'
import Delete from '@mui/icons-material/Delete';
import CircularProgress from '@mui/material/CircularProgress';
import Box from '@mui/material/Box';
import InboxIcon from '@mui/icons-material/Inbox';
import Pagination from '../Pagination/Pagination';
import { useDispatch, useSelector } from 'react-redux';
import ItemSupplier from '../Item/ItemSupplier';


function SupplierList() {
    const [pagination, setPagination] = useState({
        offset: 0,
        numberPerPage: 2,
        pageCount: 0,
    })

    const [supplierList, setSupplierList] = useState()

    const isSearchingSelector = useSelector((state) => state.supplier.isSearching)
    const supplierSelector = useSelector((state) => state.supplier.supplier)


    useEffect(() => {
        const fetchSupplier = async () => {
            try {
                const response = await fetch(' https://scm-tool.thanhpp.ninja/supplier')
                if (!response) {
                    throw new Error('somethign wrong');
                    return;
                }
                const data = await response.json()

                setPagination(prev => ({
                    ...prev,
                    pageCount: !isSearchingSelector ? Math.ceil(data.data.length / prev.numberPerPage) : Math.ceil(supplierSelector.length / prev.numberPerPage)
                }))

                const supplierDisplay = !isSearchingSelector ? data.data.slice(pagination.offset, parseInt(pagination.offset) + parseInt(pagination.numberPerPage)) : supplierSelector.slice(pagination.offset, parseInt(pagination.offset) + parseInt(pagination.numberPerPage))

                setSupplierList(supplierDisplay)
                console.log(data.data)
            } catch (err) {
                console.log('fetch wrong')
            }
        }

        fetchSupplier()
    }, [pagination.offset, pagination.numberPerPage, pagination.pageCount, isSearchingSelector, supplierSelector])

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
                <div className=' col-span-4 my-auto   truncate'>name</div>
                <div className=' col-span-4 my-auto truncate'>email</div>
                <div className=' col-span-2 my-auto truncate'>phone</div>
                <div className=' col-span-1 my-auto pl-[5px] flex justify-between truncate'>
                    <span className=' mr-[10px] rounded-sm'>Delete</span>
                    <span className=' mr-[10px] rounded-sm'>Edit</span>
                </div>
            </div>
            {/* {pagination.currentData ? pagination.currentData.map((item) => <ProductItem isSelectedAll={isSelectedAll} key={item.id} id={item.id} sku={item.sku} name={item.name} category={item.category} price={item.price} amount={item.amount} vendor={item.vendor} arrivalDate={item.arrivalDate} />) : <Box sx={{ display: 'flex' }}> */}
            {supplierList && supplierList.map(supplier => <ItemSupplier id={supplier.id} key={supplier.id} name={supplier.name} email={supplier.email} phone={supplier.phone} />)}
            {/* <CircularProgress /> */}
            {/* </Box>} */}
            <Pagination getNumberPerPage={numberPerPageHandle} pageCount={pagination.pageCount} onPageChange={handleClick} />
            {/* <Pagination /> */}
        </div>
    )
}

export default SupplierList