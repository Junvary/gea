<template>
    <div>
        <div class="page-actions">
            <el-form :inline="true" :model="searchInfo">
               {{- range .Fields}}  {{- if .FieldSearchType}} {{- if eq .FieldType "bool" }}
                <el-form-item label="{{.FieldDesc}}" prop="{{.FieldJson}}">
                <el-select v-model="searchInfo.{{.FieldJson}}" clear placeholder="请选择">
                    <el-option
                        key="true"
                        label="是"
                        value="true">
                    </el-option>
                    <el-option
                        key="false"
                        label="否"
                        value="false">
                    </el-option>
                </el-select>
                </el-form-item>
                      {{- else }}
            <el-form-item label="{{.FieldDesc}}">
                <el-input placeholder="搜索条件" v-model="searchInfo.{{.FieldJson}}"></el-input>
            </el-form-item> {{ end }} {{ end }}  {{ end }}
            <el-form-item>
                <el-button @click="onSubmit" type="primary">查询</el-button>
            </el-form-item>
            <el-form-item>
                <el-button @click="openDialog" type="primary">新增{{.Description}}</el-button>
            </el-form-item>
            <el-form-item>
                <el-popover placement="top" v-model="deleteVisible" width="160">
                    <p>确定要删除吗？</p>
                    <div style="text-align: right; margin: 0">
                        <el-button @click="deleteVisible = false" size="mini" type="text">取消</el-button>
                        <el-button @click="onDelete" size="mini" type="primary">确定</el-button>
                    </div>
                    <el-button icon="el-icon-delete" size="mini" slot="reference" type="danger">批量删除</el-button>
                </el-popover>
            </el-form-item>
            </el-form>
        </div>
        <el-table
            :data="tableData"
            @selection-change="handleSelectionChange"
            border
            ref="multipleTable"
            stripe
            style="width: 100%"
            tooltip-effect="dark"
        >
            <el-table-column type="selection" width="55"></el-table-column>
            <el-table-column label="日期" width="180">
                 <template slot-scope="scope">{{ "{{scope.row.CreatedAt|formatDate}}" }}</template>
            </el-table-column>
            {{range .Fields}}
            {{- if .DictType}}
            <el-table-column label="{{.FieldDesc}}" prop="{{.FieldJson}}" width="120">
                <template slot-scope="scope">
                    {{"{{"}}filterDict(scope.row.{{.FieldJson}},"{{.DictType}}"){{"}}"}}
                </template>
            </el-table-column>
            {{- else if eq .FieldType "bool" }}
            <el-table-column label="{{.FieldDesc}}" prop="{{.FieldJson}}" width="120">
                 <template slot-scope="scope">{{ "{{scope.row."}}{{.FieldJson}}{{"|formatBoolean}}" }}</template>
            </el-table-column>
            {{- else }}
            <el-table-column label="{{.FieldDesc}}" prop="{{.FieldJson}}" width="120"></el-table-column> {{ end }}
            {{ end }}
            <el-table-column label="按钮组">
                <template slot-scope="scope">
                    <el-button class="table-button" @click="update{{.StructName}}(scope.row)" size="small" type="primary" icon="el-icon-edit">变更</el-button>
                    <el-popover placement="top" width="160" v-model="scope.row.visible">
                        <p>确定要删除吗？</p>
                        <div style="text-align: right; margin: 0">
                              <el-button size="mini" type="text" @click="scope.row.visible = false">取消</el-button>
                              <el-button type="primary" size="mini" @click="delete{{.StructName}}(scope.row)">确定</el-button>
                        </div>
                        <el-button type="danger" icon="el-icon-delete" size="mini" slot="reference">删除</el-button>
                    </el-popover>
                </template>
            </el-table-column>
    </el-table>

    <el-pagination
        :current-page="page"
        :page-size="pageSize"
        :page-sizes="[10, 30, 50, 100]"
        :style="{float:'right',padding:'20px'}"
        :total="total"
        @current-change="handleCurrentChange"
        @size-change="handleSizeChange"
        layout="total, sizes, prev, pager, next, jumper"
    ></el-pagination>

    <el-dialog :before-close="closeDialog" :visible.sync="dialogFormVisible" title="弹窗操作">
        <el-form :model="formData" label-position="right" label-width="80px">
        {{- range .Fields}}
            <el-form-item label="{{.FieldDesc}}:">
            {{- if eq .FieldType "bool" }}
                <el-switch active-color="#13ce66" inactive-color="#ff4949" active-text="是" inactive-text="否" v-model="formData.{{.FieldJson}}" clearable ></el-switch>
            {{ end -}}
            {{- if eq .FieldType "string" }}
                <el-input v-model="formData.{{.FieldJson}}" clearable placeholder="请输入" ></el-input>
            {{ end -}}
            {{- if eq .FieldType "int" }}
            {{- if .DictType}}
                 <el-select v-model="formData.{{ .FieldJson }}" placeholder="请选择" clearable>
                     <el-option v-for="(item,key) in {{ .DictType }}Options" :key="key" :label="item.label" :value="item.value"></el-option>
                 </el-select>
            {{ else -}}
                <el-input v-model.number="formData.{{ .FieldJson }}" clearable placeholder="请输入"></el-input>
            {{ end -}}
            {{ end -}}
            {{- if eq .FieldType "time.Time" }}
                <el-date-picker type="date" placeholder="选择日期" v-model="formData.{{ .FieldJson }}" clearable></el-date-picker>
            {{ end -}}
            {{- if eq .FieldType "float64" }}
                <el-input-number v-model="formData.{{ .FieldJson }}" :precision="2" clearable></el-input-number>
            {{ end -}}
            </el-form-item>
            {{ end -}}
        </el-form>
        <div class="dialog-footer" slot="footer">
            <el-button @click="closeDialog">取 消</el-button>
            <el-button @click="enterDialog" type="primary">确 定</el-button>
        </div>
    </el-dialog>
</div>
</template>

<script>
import { create{{.StructName}}, delete{{.StructName}}, delete{{.StructName}}ByIds, update{{.StructName}}, find{{.StructName}}, get{{.StructName}}List } from "@/api/{{.PackageName}}";
import { formatTimeToStr } from "@/utils/date";
import infoList from "@/mixins/infoList";

export default {
    name: "{{.StructName}}",
    mixins: [infoList],
    data() {
        return {
            listApi: get{{ .StructName }}List,
            dialogFormVisible: false,
            visible: false,
            type: "",
            deleteVisible: false,
            multipleSelection: [],
            {{- range .Fields}}
            {{- if .DictType }}
            {{ .DictType }}
            Options:[],
            {{ end -}}
            {{end -}}

            formData: {
                {{range .Fields}}
                {{- if eq .FieldType "bool" -}}
                {{.FieldJson}}:false,
                {{ end -}}
                {{- if eq .FieldType "string" -}}
                {{.FieldJson}}:"",
                {{ end -}}
                {{- if eq .FieldType "int" -}}
                {{.FieldJson}}:0,
                {{ end -}}
                {{- if eq .FieldType "time.Time" -}}
                {{.FieldJson}}:new Date(),
                {{ end -}}
                {{- if eq .FieldType "float64" -}}
                {{.FieldJson}}:0,
                {{ end -}}
                {{ end }}
            }
        };
    },
    filters: {
        formatDate: function(time) {
            if (time != null && time != "") {
                var date = new Date(time);
                return formatTimeToStr(date, "yyyy-MM-dd hh:mm:ss");
            } else {
                return "";
            }
        },
        formatBoolean: function(bool) {
            if (bool != null) {
                return bool ? "是" :"否";
            } else {
                return "";
            }
        }
    },
    methods: {
        onSubmit() {
            this.page = 1
            this.pageSize = 10
            {{- range .Fields}} {{- if eq .FieldType "bool" }}
            if (this.searchInfo.{{.FieldJson}}==""){
              this.searchInfo.{{.FieldJson}}=null
            } {{ end }} {{ end }}
            this.getTableData()
        },
        handleSelectionChange(val) {
            this.multipleSelection = val
        },
        async onDelete() {
            const ids = []
            if(this.multipleSelection.length == 0){
                this.$message({
                    type: 'warning',
                    message: '请选择要删除的数据'
                })
                return
            }
            this.multipleSelection && this.multipleSelection.map(item => {
                ids.push(item.ID)
            })
            const res = await delete{{.StructName}}ByIds({ ids })
            if (res.code == 0) {
                this.$message({
                    type: 'success',
                    message: '删除成功'
                })
                this.deleteVisible = false
                this.getTableData()
            }
        },
        async update{{.StructName}}(row) {
            const res = await find{{.StructName}}({ ID: row.ID });
            this.type = "update";
            if (res.code == 0) {
                this.formData = res.data.re{{.Abbreviation}};
                this.dialogFormVisible = true;
            }
        },
        closeDialog() {
            this.dialogFormVisible = false;
            this.formData = {
                {{range .Fields}}
                {{- if eq .FieldType "bool" -}}
                {{.FieldJson}}:false,
                {{ end -}}
                {{- if eq .FieldType "string" -}}
                {{.FieldJson}}:"",
                {{ end -}}
                {{- if eq .FieldType "int" -}}
                {{.FieldJson}}:0,
                {{ end -}}
                {{- if eq .FieldType "time.Time" -}}
                {{.FieldJson}}:new Date(),
                {{ end -}}
                {{- if eq .FieldType "float64" -}}
                {{.FieldJson}}:0,
                {{ end -}}
                {{ end }}
            };
        },
        async delete{{.StructName}}(row) {
            this.visible = false;
            const res = await delete{{.StructName}}({ ID: row.ID });
            if (res.code == 0) {
            this.$message({
                type: "success",
                message: "删除成功"
            });
            this.getTableData();
            }
        },
        async enterDialog() {
            let res;
            switch (this.type) {
            case "create":
                res = await create{{.StructName}}(this.formData);
                break;
            case "update":
                res = await update{{.StructName}}(this.formData);
                break;
            default:
                res = await create{{.StructName}}(this.formData);
                break;
            }
            if (res.code == 0) {
            this.$message({
                type:"success",
                message:"创建/更改成功"
            })
            this.closeDialog();
            this.getTableData();
            }
        },
        openDialog() {
            this.type = "create";
            this.dialogFormVisible = true;
        }
    },
    async created() {
        await this.getTableData();
        {{ range .Fields -}}
        {{- if .DictType }}
        await this.getDict("{{.DictType}}");
        {{ end -}}
        {{- end }}
    }
};
</script>

<style lang="scss" scopde>
</style>