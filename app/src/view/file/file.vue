<template>
    <div >
        <Table :columns="columns1" :data="data1"></Table>
    </div>
</template>

<script>
    export default {
        data () {
            return {
                columns1: [
                    {
                        title: '文件名',
                        key: 'name'
                    },
                    {
                        title: '大小',
                        key: 'size'
                    }
                ],
                data1: [
                    {
                        name: '4221678333534_.pic.jpg',
                        size: 221255,
                    },
                ]
            }
        }
    }
</script>