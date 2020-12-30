<template>
    <div>
        <el-form :inline="true" :model="searchInfo" class="demo-form-inline">
            <el-form-item label="请求方法">
                <el-input placeholder="搜索条件" v-model="searchInfo.method"></el-input>
            </el-form-item>
            <el-form-item label="请求路径">
                <el-input placeholder="搜索条件" v-model="searchInfo.path"></el-input>
            </el-form-item>
            <el-form-item label="结果状态码">
                <el-input placeholder="搜索条件" v-model="searchInfo.status"></el-input>
            </el-form-item>
            <el-form-item>
                <el-button @click="onSubmit" type="primary">查询</el-button>
            </el-form-item>
            <el-form-item>
                <el-button @click="onDelete" type="danger">批量删除</el-button>
            </el-form-item>
        </el-form>
        <el-table :data="tableData" @selection-change="handleSelectionChange" border ref="multipleTable" stripe
            style="width: 100%" tooltip-effect="dark" v-loading="loading">
            <el-table-column type="selection" min-width="55"></el-table-column>
            <el-table-column label="操作人" min-width="140">
                <template slot-scope="scope">
                    <div>{{scope.row.user.userName}}({{scope.row.user.nickName}})</div>
                </template>
            </el-table-column>
            <el-table-column label="日期" min-width="180">
                <template slot-scope="scope">{{scope.row.CreatedAt|formatDate}}</template>
            </el-table-column>
            <el-table-column label="状态码" prop="status" min-width="80">
                <template slot-scope="scope">
                    <div>
                        <el-tag type="success">{{ scope.row.status }}</el-tag>
                    </div>
                </template>
            </el-table-column>
            <el-table-column label="请求ip" prop="ip" min-width="120"></el-table-column>
            <el-table-column label="请求方法" prop="method" min-width="100"></el-table-column>
            <el-table-column label="请求路径" prop="path" min-width="300"></el-table-column>
            <el-table-column label="请求" prop="path" min-width="60">
                <template slot-scope="scope">
                    <div>
                        <el-popover placement="top-start" trigger="hover" v-if="scope.row.body">
                            <div class="popover-box">
                                <pre>{{fmtBody(scope.row.body)}}</pre>
                            </div>
                            <i class="el-icon-view" slot="reference"></i>
                        </el-popover>

                        <span v-else>无</span>
                    </div>
                </template>
            </el-table-column>
            <el-table-column label="响应" prop="path" min-width="60">
                <template slot-scope="scope">
                    <div>
                        <el-popover placement="top-start" trigger="hover" v-if="scope.row.resp">
                            <div class="popover-box">
                                <pre>{{fmtBody(scope.row.resp)}}</pre>
                            </div>
                            <i class="el-icon-view" slot="reference"></i>
                        </el-popover>
                        <span v-else>无</span>
                    </div>
                </template>
            </el-table-column>
            <el-table-column label="按钮组" width="200">
                <template slot-scope="scope">
                    <el-button size="mini" @click="deleteLogOperation(scope.row)" type="danger">删除</el-button>
                </template>
            </el-table-column>
        </el-table>
        <el-pagination :current-page="page" :page-size="pageSize" :page-sizes="[10, 30, 50, 100]"
            :style="{float:'right',padding:'20px'}" :total="total" @current-change="handleCurrentChange"
            @size-change="handleSizeChange" layout="total, sizes, prev, pager, next, jumper"></el-pagination>
    </div>
</template>

<script>
import { deleteLogOperation, getLogOperationList, deleteLogOperationByIds } from '@/api/log' //  此处请自行替换地址
import { formatTimeToStr } from '@/utils/date'
import infoList from '@/mixins/infoList'

export default {
    name: 'Operation',
    mixins: [infoList],
    data() {
        return {
            listApi: getLogOperationList,
            dialogFormVisible: false,
            type: '',
            multipleSelection: [],
            formData: {
                ip: null,
                method: null,
                path: null,
                status: null,
                latency: null,
                agent: null,
                error_message: null,
                user_id: null,
            },
        }
    },
    filters: {
        formatDate: function (time) {
            if (time != null && time != '') {
                var date = new Date(time)
                return formatTimeToStr(date, 'yyyy-MM-dd hh:mm:ss')
            } else {
                return ''
            }
        },
        formatBoolean: function (bool) {
            if (bool != null) {
                return bool ? '是' : '否'
            } else {
                return ''
            }
        },
    },
    methods: {
        //条件搜索前端看此方法
        onSubmit() {
            this.page = 1
            this.pageSize = 10
            this.getTableData()
        },
        handleSelectionChange(val) {
            this.multipleSelection = val
        },
        async onDelete() {
            this.$confirm('此操作将删除所选日志, 是否继续?', '提示', {
                confirmButtonText: '确定',
                cancelButtonText: '取消',
                type: 'warning',
                lockScroll: false,
                showClose: false,
                closeOnClickModal: false,
                closeOnPressEscape: false,
            })
                .then(async () => {
                    const ids = []
                    this.multipleSelection &&
                        this.multipleSelection.map((item) => {
                            ids.push(item.ID)
                        })
                    const res = await deleteLogOperationByIds({ ids })
                    if (res.code == 0) {
                        this.$message({
                            type: 'success',
                            message: '删除成功',
                        })
                        if (this.tableData.length == 1) {
                            this.page--
                        }
                        this.getTableData()
                    }
                })
                .catch(() => {
                    this.$message({
                        type: 'info',
                        message: '已取消删除',
                    })
                })
        },
        async deleteLogOperation(row) {
            this.$confirm('此操作将删除该条日志, 是否继续?', '提示', {
                confirmButtonText: '确定',
                cancelButtonText: '取消',
                type: 'warning',
                lockScroll: false,
                showClose: false,
                closeOnClickModal: false,
                closeOnPressEscape: false,
            })
                .then(async () => {
                    const res = await deleteLogOperation({ ID: row.ID })
                    if (res.code == 0) {
                        this.$message({
                            type: 'success',
                            message: '删除成功',
                        })
                        if (this.tableData.length == 1) {
                            this.page--
                        }
                        this.getTableData()
                    }
                })
                .catch(() => {
                    this.$message({
                        type: 'info',
                        message: '已取消删除',
                    })
                })
        },
        fmtBody(value) {
            try {
                return JSON.parse(value)
            } catch (err) {
                return value
            }
        },
    },
    created() {
        this.getTableData()
    },
}
</script>
