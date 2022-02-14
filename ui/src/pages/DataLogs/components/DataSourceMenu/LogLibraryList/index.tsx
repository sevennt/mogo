import logLibraryListStyles from "@/pages/DataLogs/components/DataSourceMenu/LogLibraryList/index.less";
import { Empty, Spin } from "antd";
import { useModel } from "@@/plugin-model/useModel";
import { useIntl } from "umi";
import LogLibraryItem from "@/pages/DataLogs/components/DataSourceMenu/LogLibraryList/LogLibraryItem";
import DatabaseViewsDraw from "@/pages/DataLogs/components/DataSourceMenu/LogLibraryList/DatabaseViewsDraw";
import { useState } from "react";
import { TablesResponse } from "@/services/dataLogs";
import LogLibraryInfoDraw from "@/pages/DataLogs/components/DataSourceMenu/LogLibraryList/LogLibraryInfoDraw";

type LogLibraryListProps = {
  list: TablesResponse[];
};

const LogLibraryList = (props: LogLibraryListProps) => {
  const { list } = props;
  const { getLogLibraries } = useModel("dataLogs");

  const [selectedLogLibrary, setSelectedLogLibrary] = useState<
    TablesResponse | undefined
  >();
  const onChangeSelected = (logLibrary: TablesResponse) => {
    setSelectedLogLibrary(logLibrary);
  };

  const i18n = useIntl();

  return (
    <div className={logLibraryListStyles.logLibraryListMain}>
      <Spin
        spinning={getLogLibraries.loading}
        tip={i18n.formatMessage({ id: "spin" })}
      >
        {list.length > 0 ? (
          <ul>
            {list.map((item, index) => (
              <LogLibraryItem
                onChange={onChangeSelected}
                logLibrary={item}
                key={index}
              />
            ))}
          </ul>
        ) : (
          <Empty
            image={Empty.PRESENTED_IMAGE_SIMPLE}
            description={i18n.formatMessage({
              id: "datasource.logLibrary.empty",
            })}
          />
        )}
      </Spin>
      <DatabaseViewsDraw logLibrary={selectedLogLibrary as TablesResponse} />
      <LogLibraryInfoDraw logLibrary={selectedLogLibrary as TablesResponse} />
    </div>
  );
};

export default LogLibraryList;
