import React, { useEffect, useState } from 'react'
import Delete from '@mui/icons-material/Delete';
import CircularProgress from '@mui/material/CircularProgress';
import Box from '@mui/material/Box';
import WarehouseIcon from '@mui/icons-material/Warehouse';
import ItemWarehouse from '../Item/ItemWarehouse';
import Pagination from '../Pagination/Pagination';
import { useSelector } from 'react-redux';


function WarehouseList() {
    const [pagination, setPagination] = useState({
        offset: 0,
        numberPerPage: 2,
        pageCount: 0,
    })

    const [warehouseList, setWarehouseList] = useState()

    const isSearchingSelector = useSelector((state) => state.warehouse.isSearching)
    const warehouseSelector = useSelector((state) => state.warehouse.warehouse)


    useEffect(() => {
        const fetchWarehouse = async () => {
            console.log(warehouseSelector)
            try {
                const response = await fetch(`https://scm-tool.thanhpp.ninja/storage`)
                if (!response) {
                    throw new Error('something wrong');
                    return;
                }
                const data = await response.json()

                setPagination(prev => ({
                    ...prev,
                    pageCount: !isSearchingSelector ? Math.ceil(data.data.length / prev.numberPerPage) : Math.ceil(warehouseSelector.length / prev.numberPerPage)
                }))

                const warehousesDisplay = !isSearchingSelector ? data.data.slice(pagination.offset, parseInt(pagination.offset) + parseInt(pagination.numberPerPage)) : warehouseSelector.slice(pagination.offset, parseInt(pagination.offset) + parseInt(pagination.numberPerPage))

                setWarehouseList(warehousesDisplay)


            } catch (err) {
                console.log('fetch wrong')
            }
        }

        fetchWarehouse()
    }, [pagination.offset, pagination.numberPerPage, pagination.pageCount, isSearchingSelector, warehouseSelector])


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
                    <WarehouseIcon className='mr-[10px] leading-[24px] ml-[20px]' />
                </div>
                <div className=' col-span-3 my-auto   truncate'>name</div>
                <div className=' col-span-5 my-auto truncate'>description</div>
                <div className=' col-span-2 my-auto truncate'>location</div>
                <div className=' col-span-1 my-auto pl-[5px] flex justify-between truncate'>
                    <span className=' mr-[10px] rounded-sm'>Delete</span>
                    <span className=' mr-[10px] rounded-sm'>Edit</span>
                </div>
            </div>
            {warehouseList && warehouseList.map(warehouse => <ItemWarehouse id={warehouse.id} key={warehouse.id} name={warehouse.name} desc={warehouse.desc} location={warehouse.location} />)}
            {/* <CircularProgress /> */}
            {/* </Box>} */}
            <Pagination getNumberPerPage={numberPerPageHandle} pageCount={pagination.pageCount} onPageChange={handleClick} />
        </div>
    )
}

export default WarehouseList