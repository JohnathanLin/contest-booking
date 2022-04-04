<template>
  <div class="hello">
    <h1>赛事预约提醒系统</h1>
    <el-card class="box-card" style="margin-bottom: 15px;">
      <div slot="header" class="clearfix" style="position: relative;">
        <span>近期比赛</span>
        <el-button style="float: right; display:inline-block; padding: 3px 0" type="text" size="small">
          <a href="https://lpl.qq.com/es/schedule.shtml" target="_blank">访问lpl赛程页</a></el-button>
        <el-button style="float: right; display:inline-block; padding: 3px 0" type="text" @click="fetchLPLContestData">刷新</el-button>
      </div>
      <el-table
        :data="contestList"
        v-loading="contestListLoading"
        style="width: 100%">
        <el-table-column
          prop="contestId"
          label="比赛id"
          width="80">
        </el-table-column>
        <el-table-column
          prop="status"
          label="状态"
          width="100">
          <template slot-scope="scope">
            <el-tag type="info" v-if="scope.row.status === 1">未开始</el-tag>
            <el-tag type="warning" v-if="scope.row.status === 2">进行中</el-tag>
            <el-tag type="success" v-if="scope.row.status === 3">已结束</el-tag>
          </template>
        </el-table-column>
        <el-table-column
          prop="contestName"
          label="赛事名称">
        </el-table-column>
        <el-table-column
          prop="contestStage"
          label="赛事阶段">
          <template slot-scope="scope">
            {{ scope.row.contestStage }}  {{ scope.row.contestMarchName }}
          </template>
        </el-table-column>
        <el-table-column
          prop="time"
          label="比赛时间"
        >
        </el-table-column>
        <el-table-column
          prop="result"
          label="对阵情况">
        </el-table-column>

        <el-table-column
          prop="operation"
          label="操作">
          <template slot-scope="scope">
            <el-button type="text" size="small" @click="showBookingDialog(scope.row)">预约</el-button>
          </template>
        </el-table-column>
      </el-table>
    </el-card>
    <el-card class="box-card">
      <div slot="header" class="clearfix">
        <span>当前预约</span>
        <el-button style="float: right; padding: 3px 0" type="text" @click="fetchBookingData">刷新</el-button>
      </div>
      <el-table
        :data="bookingList"
        v-loading="bookingListLoading"
        style="width: 100%">
        <el-table-column
          prop="bookingId"
          label="预约id"
          width="80">
        </el-table-column>

        <el-table-column
          prop="contestId"
          label="比赛id"
          width="80">
        </el-table-column>
        <el-table-column
          prop="status"
          label="状态"
          width="80">
          <template slot-scope="scope">
            <el-tag type="info" v-if="scope.row.bookingStatus === 1">未开始</el-tag>
            <el-tag type="warning" v-if="scope.row.bookingStatus === 2">进行中</el-tag>
            <el-tag type="success" v-if="scope.row.bookingStatus === 3">预约成功</el-tag>
            <el-tag type="danger" v-if="scope.row.contestStatus === 4">预约失败</el-tag>
          </template>
        </el-table-column>
        <el-table-column
          prop="contestName"
          label="比赛名称"
        >
        </el-table-column>
        <el-table-column
          prop="contestTime"
          label="比赛时间"
        >
        </el-table-column>
        <el-table-column
          prop="target"
          label="预约目标比分">
          <template slot-scope="scope">
            {{ scope.row.teamA }} {{ scope.row.bookingScoreA }} : {{ scope.row.bookingScoreB }} {{ scope.row.teamB }}
          </template>
        </el-table-column>
        <el-table-column
          prop="result"
          label="结果">
          <template slot-scope="scope">
            {{ scope.row.teamA }} {{ scope.row.resultScoreA }} : {{ scope.row.resultScoreB }} {{ scope.row.teamB }}
          </template>
        </el-table-column>
        <el-table-column
          prop="bookingEmail"
          label="预约邮箱"
        >
        </el-table-column>
        <el-table-column
          prop="operation"
          label="操作">
          <template slot-scope="scope">
            <el-button type="text" size="small"  @click="cancelBooking(scope.row)">取消并删除</el-button>
          </template>
        </el-table-column>
      </el-table>
    </el-card>

    <el-dialog :visible.sync="visible" title="预约比赛" :before-close="handleBookingDialogClose">
<!--      <p>您要预约的比赛为: 2022职业联赛 春季赛季后赛 TES VS JDG</p>-->
      <p>您要预约的比赛为: {{ bookingFormData.bookingContestBrief }}</p>
      {{ bookingFormData.contestMarchName.split("vs")[0] }}:  <el-input-number v-model="bookingFormData.targetScoreA" :min="0" :max="3" label="描述文字"></el-input-number>
      {{ bookingFormData.contestMarchName.split("vs")[1] }}:  <el-input-number v-model="bookingFormData.targetScoreB" :min="0" :max="3" label="描述文字"></el-input-number>

      <el-input
        type="textarea"
        :rows="2"
        placeholder="预约邮箱,多个邮箱请用用英文分号;分开"
        v-model="bookingFormData.bookingEmail">
      </el-input>
      <el-button type="text" size="small" @click="startBooking">预约</el-button>
    </el-dialog>
  </div>
</template>

<script>
  // import request from '@/utils/request'
  import axios from 'axios'
export default {
  name: 'HelloWorld',
  data () {
    return {
      visible: false,
      contestList: [],
      contestListLoading: false,
      bookingListLoading: false,
      bookingList: [],


      bookingFormData: {
        contestId: 0,
        bookingEmail: '',
        bookingContestBrief: '',
        contestMarchName: 'a vs b',
        bookingContestId: 0,
        targetScoreA: 0,
        targetScoreB: 0,
        contestTime: null,
      }

    }
  },
  created() {
    this.fetchLPLContestData()
    this.fetchBookingData()
  },
  methods: {
    fetchLPLContestData() {
      this.contestListLoading = true
      axios.get("/getContestList")
        .then(response => {
          this.contestList = response.data
          this.contestListLoading = false
          // console.log(response);
        })
        .catch(function (error) {
          console.log(error);
        });
    },

    fetchBookingData() {
      this.bookingListLoading = true
      axios.get("/getBookingList")
        .then(response => {
          this.bookingList = response.data
          this.bookingListLoading = false
        })
    },


    showBookingDialog(row) {
      this.visible = true
      this.bookingFormData.bookingContestBrief = row.contestName + " " + row.contestStage + " " + row.contestMarchName
      this.bookingFormData.contestMarchName = row.contestMarchName
      this.bookingFormData.bookingContestId = row.contestId
      this.bookingFormData.targetScoreA = 0
      this.bookingFormData.targetScoreB = 0
      this.bookingFormData.contestTime = row.time
      console.log("contestTime" + this.contestTime)
      this.bookingFormData.contestId = row.contestId
    },
    handleBookingDialogClose(done) {
      this.$confirm('确认关闭？')
        .then(_ => {
          done();
        })
        .catch(_ => {});
    },

    startBooking() {
      this.$confirm('确认提交预约？')
        .then(_ => {
          var jsonObj = {
            "contestId": this.bookingFormData.contestId,
            // "bookingStatus": 1,
            "contestTime": this.bookingFormData.contestTime,
            "teamA": this.bookingFormData.contestMarchName.split("vs")[0],
            "teamB":  this.bookingFormData.contestMarchName.split("vs")[1],
            "bookingScoreA": this.bookingFormData.targetScoreA,
            "bookingScoreB": this.bookingFormData.targetScoreB,
            // "resultScoreA": 0,
            // "resultScoreB": 0,
            "bookingEmail": this.bookingFormData.bookingEmail,
            "contestName": this.bookingFormData.bookingContestBrief
          }
          axios.post("/startBooking", JSON.stringify(jsonObj)) .then(response => {
            this.$message.success("预约成功")
            this.visible = false
            this.fetchBookingData()
          })
        })
        .catch(_ => {});

    },
    cancelBooking(row) {
      this.$confirm('是否删除预约' + row.bookingId + '？')
        .then(_ => {

          axios.delete("/cancelBooking?bookingId=" + row.bookingId) .then(response => {
            this.$message.success("删除成功")
            // this.visible = false
            this.fetchBookingData()
          })
        })
        .catch(_ => {});

    }
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
h1, h2 {
  font-weight: normal;
}
ul {
  list-style-type: none;
  padding: 0;
}
li {
  display: inline-block;
  margin: 0 10px;
}
a {
  color: #42b983;
}
</style>
