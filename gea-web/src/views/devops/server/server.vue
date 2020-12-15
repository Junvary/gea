<template>
    <div>
        <el-form :inline="true" :model="searchInfo">
            <el-form-item label="区域">
                <el-input placeholder="搜索条件" v-model="searchInfo.area"></el-input>
            </el-form-item>
            <el-form-item label="位置">
                <el-input placeholder="搜索条件" v-model="searchInfo.location"></el-input>
            </el-form-item>
            <el-form-item label="类别">
                <el-input placeholder="搜索条件" v-model="searchInfo.category"></el-input>
            </el-form-item>
            <el-form-item label="操作系统">
                <el-input placeholder="搜索条件" v-model="searchInfo.os"></el-input>
            </el-form-item>
            <el-form-item label="名称">
                <el-input placeholder="搜索条件" v-model="searchInfo.name"></el-input>
            </el-form-item>
            <el-form-item label="IP地址">
                <el-input placeholder="搜索条件" v-model="searchInfo.ip"></el-input>
            </el-form-item>
            <el-form-item label="端口">
                <el-input placeholder="搜索条件" v-model="searchInfo.port"></el-input>
            </el-form-item>
            <el-form-item label="描述备注">
                <el-input placeholder="搜索条件" v-model="searchInfo.remark"></el-input>
            </el-form-item>
            <el-form-item>
                <el-button @click="onSubmit" type="primary">查询</el-button>
            </el-form-item>
            <el-form-item>
                <el-button @click="openDialog" type="primary">新增远程主机</el-button>
            </el-form-item>
            <el-form-item>
                <el-popover placement="top" v-model="deleteVisible" width="160">
                    <p>确定要删除吗？</p>
                    <div style="text-align: right; margin: 0">
                        <el-button @click="deleteVisible = false" size="mini" type="text">取消</el-button>
                        <el-button @click="onDelete" size="mini" type="primary">确定</el-button>
                    </div>
                    <el-button slot="reference" type="danger">批量删除</el-button>
                </el-popover>
            </el-form-item>
        </el-form>
        <el-table :data="tableData" @selection-change="handleSelectionChange" border ref="multipleTable" stripe
            style="width: 100%" tooltip-effect="dark">
            <el-table-column type="selection" width="55"></el-table-column>
            <el-table-column label="日期" width="180">
                <template slot-scope="scope">{{scope.row.CreatedAt|formatDate}}</template>
            </el-table-column>

            <el-table-column label="区域" prop="area" width="120"></el-table-column>

            <el-table-column label="位置" prop="location" width="120"></el-table-column>

            <el-table-column label="类别" prop="category" width="120"></el-table-column>

            <el-table-column label="操作系统" prop="os" width="120">
                <template slot-scope="scope">
                    {{filterDict(scope.row.os,"os")}}
                </template>
            </el-table-column>

            <el-table-column label="名称" prop="name" width="120"></el-table-column>

            <el-table-column label="IP地址" prop="ip" width="120"></el-table-column>

            <el-table-column label="端口" prop="port" width="120"></el-table-column>

            <el-table-column label="描述备注" prop="remark" width="120"></el-table-column>

            <el-table-column label="按钮组">
                <template slot-scope="scope">
                    <el-button class="table-button" @click="updateDevopsServer(scope.row)" size="small" type="primary">
                        变更
                    </el-button>
                    <el-button @click="deleteDevopsServer(scope.row)" type="danger" size="small">删除</el-button>
                </template>
            </el-table-column>
        </el-table>

        <el-pagination :current-page="page" :page-size="pageSize" :page-sizes="[10, 30, 50, 100]"
            :style="{float:'right',padding:'20px'}" :total="total" @current-change="handleCurrentChange"
            @size-change="handleSizeChange" layout="total, sizes, prev, pager, next, jumper"></el-pagination>

        <el-dialog :before-close="closeDialog" :visible.sync="dialogFormVisible" title="弹窗操作">
            <el-form ref="geaForm" :model="formData" :rules="rules" label-position="right" label-width="100px">
                <el-form-item label="区域:" prop="area">
                    <el-input v-model="formData.area" clearable placeholder="请输入"></el-input>
                </el-form-item>

                <el-form-item label="位置:" prop="location">
                    <el-input v-model="formData.location" clearable placeholder="请输入"></el-input>
                </el-form-item>

                <el-form-item label="类别:" prop="category">
                    <el-input v-model="formData.category" clearable placeholder="请输入"></el-input>
                </el-form-item>

                <el-form-item label="操作系统:" prop="os">
                    <el-select v-model="formData.os" placeholder="请选择">
                        <el-option v-for="(item, key) in osOptions" :key="key" :label="item.label" :value="item.value">
                        </el-option>
                    </el-select>
                </el-form-item>
                <el-form-item label="名称:" prop="name">
                    <el-input v-model="formData.name" clearable placeholder="请输入"></el-input>
                </el-form-item>

                <el-form-item label="IP地址:" prop="ip">
                    <el-input v-model="formData.ip" clearable placeholder="请输入"></el-input>
                </el-form-item>

                <el-form-item label="端口:" prop="port">
                    <el-input v-model="formData.port" clearable placeholder="请输入"></el-input>
                </el-form-item>

                <el-form-item label="描述备注:" prop="remark">
                    <el-input v-model="formData.remark" clearable placeholder="请输入"></el-input>
                </el-form-item>
            </el-form>
            <div class="dialog-footer" slot="footer">
                <el-button @click="closeDialog">取 消</el-button>
                <el-button @click="enterDialog" type="primary">确 定</el-button>
            </div>
        </el-dialog>
    </div>
</template>

<script>
import { createDevopsServer, deleteDevopsServer, deleteDevopsServerByIds, updateDevopsServer, findDevopsServer, getDevopsServerList } from '@/api/devops_server'
import { formatTimeToStr } from '@/utils/date'
import infoList from '@/mixins/infoList'

export default {
    name: 'DevopsServer',
    mixins: [infoList],
    data() {
        return {
            listApi: getDevopsServerList,
            dialogFormVisible: false,
            visible: false,
            type: '',
            deleteVisible: false,
            multipleSelection: [],
            osOptions: [],
            formData: {
                area: '',
                location: '',
                category: '',
                os: '',
                name: '',
                ip: '',
                port: '',
                remark: '',
            },
            rules: {
                area: [
                    {
                        required: true,
                        message: '请输入区域',
                        trigger: 'blur',
                    },
                ],
                location: [
                    {
                        required: true,
                        message: '请输入位置',
                        trigger: 'blur',
                    },
                ],
                category: [
                    {
                        required: true,
                        message: '请输入类别',
                        trigger: 'blur',
                    },
                ],
                os: [
                    {
                        required: true,
                        message: '请输入操作系统',
                        trigger: 'blur',
                    },
                ],
                name: [
                    {
                        required: true,
                        message: '请输入名称',
                        trigger: 'blur',
                    },
                ],
                ip: [
                    {
                        required: true,
                        message: '请输入IP地址',
                        trigger: 'blur',
                    },
                ],
                port: [
                    {
                        required: true,
                        message: '请输入端口',
                        trigger: 'blur',
                    },
                ],
                remark: [{}],
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
        onSubmit() {
            this.page = 1
            this.pageSize = 10
            this.getTableData()
        },
        handleSelectionChange(val) {
            this.multipleSelection = val
        },
        async onDelete() {
            const ids = []
            if (this.multipleSelection.length == 0) {
                this.$message({
                    type: 'warning',
                    message: '请选择要删除的数据',
                })
                return
            }
            this.multipleSelection &&
                this.multipleSelection.map((item) => {
                    ids.push(item.ID)
                })
            const res = await deleteDevopsServerByIds({ ids })
            if (res.code == 0) {
                this.$message({
                    type: 'success',
                    message: '删除成功',
                })
                this.deleteVisible = false
                this.getTableData()
            }
        },
        async updateDevopsServer(row) {
            const res = await findDevopsServer({ ID: row.ID })
            this.type = 'update'
            if (res.code == 0) {
                this.formData = res.data.reserver
                this.dialogFormVisible = true
            }
        },
        closeDialog() {
            this.dialogFormVisible = false
            this.formData = {
                area: '',
                location: '',
                category: '',
                os: 0,
                name: '',
                ip: '',
                port: '',
                remark: '',
            }
        },
        async deleteDevopsServer(row) {
            this.$confirm('此操作将永久删除该主机, 是否继续?', '提示', {
                confirmButtonText: '确定',
                cancelButtonText: '取消',
                type: 'warning',
                lockScroll: false,
                showClose: false,
                closeOnClickModal: false,
                closeOnPressEscape: false,
            })
                .then(async () => {
                    this.visible = false
                    const res = await deleteDevopsServer({ ID: row.ID })
                    if (res.code == 0) {
                        this.$message({
                            type: 'success',
                            message: '删除成功',
                        })
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
        async enterDialog() {
            this.$refs['geaForm'].validate(async (valid) => {
                if (!valid) return
                let res
                switch (this.type) {
                    case 'create':
                        res = await createDevopsServer(this.formData)
                        break
                    case 'update':
                        res = await updateDevopsServer(this.formData)
                        break
                    default:
                        res = await createDevopsServer(this.formData)
                        break
                }
                if (res.code == 0) {
                    this.$message({
                        type: 'success',
                        message: '创建/更改成功',
                    })
                    this.closeDialog()
                    this.getTableData()
                }
            })
        },
        openDialog() {
            this.type = 'create'
            this.dialogFormVisible = true
        },
    },
    async created() {
        await this.getTableData()
        await this.getDict('os')
    },
}
</script>

<style lang="scss" scopde>
</style>