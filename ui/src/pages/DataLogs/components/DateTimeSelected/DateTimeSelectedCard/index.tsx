import darkTimeStyles from "@/pages/DataLogs/components/DateTimeSelected/index.less";
import { Tabs } from "antd";
import CustomTimeInterval from "@/pages/DataLogs/components/DateTimeSelected/CustomTimeInterval";
import RelativeTime from "@/pages/DataLogs/components/DateTimeSelected/RelativeTime";
import { useModel } from "@@/plugin-model/useModel";
import { TimeRangeType } from "@/config/config";
import { DarkTimeContext } from "@/pages/DataLogs/components/DateTimeSelected";
import { useContext } from "react";
const { TabPane } = Tabs;
type DarkTimeSelectedCardProps = {};
const DateTimeSelectedCard = (props: DarkTimeSelectedCardProps) => {
  const { activeTabKey, onChangeActiveTabKey } = useModel("dataLogs");
  const { TabName } = useContext(DarkTimeContext);
  return (
    <div className={darkTimeStyles.darkTimeSelectCard}>
      <Tabs
        tabBarStyle={{ padding: 0 }}
        activeKey={activeTabKey}
        size="small"
        onTabClick={onChangeActiveTabKey}
        defaultActiveKey={activeTabKey}
      >
        <TabPane
          forceRender
          tab={TabName[TimeRangeType.Relative]}
          key={TimeRangeType.Relative}
        >
          <RelativeTime />
        </TabPane>
        <TabPane
          forceRender
          tab={TabName[TimeRangeType.Custom]}
          key={TimeRangeType.Custom}
        >
          <CustomTimeInterval />
        </TabPane>
      </Tabs>
    </div>
  );
};
export default DateTimeSelectedCard;
