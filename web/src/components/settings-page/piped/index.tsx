import {
  Box,
  Button,
  Dialog,
  DialogContent,
  DialogTitle,
  Divider,
  Paper,
  Table,
  TableBody,
  TableCell,
  TableContainer,
  TableHead,
  TableRow,
  Toolbar,
} from "@mui/material";
import {
  Add as AddIcon,
  Close as CloseIcon,
  FilterList as FilterIcon,
  Update as UpgradeIcon,
} from "@mui/icons-material";
import Alert from "@mui/material/Alert";
import { createSelector } from "@reduxjs/toolkit";
import { FC, memo, useCallback, useEffect, useState } from "react";
import { TextWithCopyButton } from "~/components/text-with-copy-button";
import {
  UI_TEXT_ADD,
  UI_TEXT_CLOSE,
  UI_TEXT_FILTER,
  UI_TEXT_HIDE_FILTER,
  UI_TEXT_UPGRADE,
} from "~/constants/ui-text";
import { REQUEST_PIPED_RESTART_SUCCESS } from "~/constants/toast-text";
import { useAppDispatch, useAppSelector } from "~/hooks/redux";
import { useInterval } from "~/hooks/use-interval";
import {
  clearRegisteredPipedInfo,
  disablePiped,
  enablePiped,
  restartPiped,
  fetchPipeds,
  fetchReleasedVersions,
  fetchBreakingChanges,
  Piped,
  RegisteredPiped,
  selectAllPipeds,
} from "~/modules/pipeds";
import { addToast } from "~/modules/toasts";
import { AppState } from "~/store";
import { AddPipedDialog } from "./components/add-piped-dialog";
import { EditPipedDialog } from "./components/edit-piped-dialog";
import { FilterValues, PipedFilter } from "./components/piped-filter";
import { PipedTableRow } from "./components/piped-table-row";
import { UpgradePipedDialog } from "./components/upgrade-dialog";
import { TableCellNoWrap } from "../styles";
import { useGetProject } from "~/queries/project/use-get-project";

const filterValue = (
  _: AppState,
  filterValue: FilterValues
): boolean | undefined => filterValue.enabled;

const selectFilteredPipeds = createSelector(
  [selectAllPipeds, filterValue],
  (pipeds, enabled) => {
    switch (enabled) {
      case true:
        return pipeds.filter((piped) => piped.disabled === false);
      case false:
        return pipeds.filter((piped) => piped.disabled);
      default:
        return pipeds;
    }
  }
);

const OLD_KEY_ALERT_MESSAGE =
  "The old key is still there.\nDo not forget to delete it once you update your Piped to use this new key.";

const FETCH_INTERVAL = 30000;

export const SettingsPipedPage: FC = memo(function SettingsPipedPage() {
  const [openFilter, setOpenFilter] = useState(false);
  const [isOpenForm, setIsOpenForm] = useState(false);
  const [editPipedId, setEditPipedId] = useState<string | null>(null);
  const [filterValues, setFilterValues] = useState<FilterValues>({
    enabled: true,
  });
  const dispatch = useAppDispatch();
  const { data: projectDetail } = useGetProject();
  const pipeds = useAppSelector<Piped.AsObject[]>((state) =>
    selectFilteredPipeds(state, filterValues)
  );

  useEffect(() => {
    dispatch(fetchReleasedVersions());
  }, [dispatch]);

  useEffect(() => {
    if (projectDetail?.id) {
      dispatch(fetchBreakingChanges({ projectId: projectDetail.id }));
    }
  }, [dispatch, projectDetail?.id]);

  const releasedVersions = useAppSelector<string[]>(
    (state) => state.pipeds.releasedVersions
  );

  const breakingChangesNote = useAppSelector<string | null>(
    (state) => state.pipeds.breakingChangesNote
  );
  // TODO: Remove this console.log
  console.log("[DEBUG]", breakingChangesNote);

  const [isUpgradeDialogOpen, setIsUpgradeDialogOpen] = useState(false);
  const handleUpgradeDialogClose = (): void => setIsUpgradeDialogOpen(false);

  const registeredPiped = useAppSelector<RegisteredPiped | null>(
    (state) => state.pipeds.registeredPiped
  );

  const handleDisable = useCallback(
    (id: string) => {
      dispatch(disablePiped({ pipedId: id })).then(() => {
        dispatch(fetchPipeds(true));
      });
    },
    [dispatch]
  );
  const handleEnable = useCallback(
    (id: string) => {
      dispatch(enablePiped({ pipedId: id })).then(() => {
        dispatch(fetchPipeds(true));
      });
    },
    [dispatch]
  );
  const handleRestart = useCallback(
    (id: string) => {
      dispatch(restartPiped({ pipedId: id })).then(() => {
        dispatch(
          addToast({
            message: REQUEST_PIPED_RESTART_SUCCESS,
            severity: "success",
          })
        );
      });
    },
    [dispatch]
  );

  const handleEdit = useCallback((id: string) => {
    setEditPipedId(id);
  }, []);

  const handleClose = useCallback(() => {
    setIsOpenForm(false);
  }, []);

  const handleClosePipedInfo = useCallback(() => {
    dispatch(clearRegisteredPipedInfo());
    dispatch(fetchPipeds(true));
  }, [dispatch]);

  const handleEditClose = useCallback(() => {
    setEditPipedId(null);
  }, []);

  useInterval(() => {
    dispatch(fetchPipeds(true));
  }, FETCH_INTERVAL);

  return (
    <>
      <Toolbar variant="dense">
        <Button
          color="primary"
          startIcon={<AddIcon />}
          onClick={() => setIsOpenForm(true)}
        >
          {UI_TEXT_ADD}
        </Button>
        <Box
          sx={{
            flexGrow: 1,
          }}
        />
        <Button
          color="primary"
          startIcon={<UpgradeIcon />}
          onClick={() => setIsUpgradeDialogOpen(true)}
        >
          {UI_TEXT_UPGRADE}
        </Button>
        <Button
          color="primary"
          startIcon={openFilter ? <CloseIcon /> : <FilterIcon />}
          onClick={() => setOpenFilter(!openFilter)}
        >
          {openFilter ? UI_TEXT_HIDE_FILTER : UI_TEXT_FILTER}
        </Button>
      </Toolbar>
      <Divider />
      <Box
        sx={{
          display: "flex",
          flex: 1,
          overflow: "hidden",
        }}
      >
        <TableContainer component={Paper} square>
          <Table aria-label="piped list" size="small" stickyHeader>
            <TableHead>
              <TableRow>
                <TableCellNoWrap>Name</TableCellNoWrap>
                <TableCellNoWrap>ID</TableCellNoWrap>
                <TableCellNoWrap>Version</TableCellNoWrap>
                <TableCellNoWrap>Description</TableCellNoWrap>
                <TableCellNoWrap>Started At</TableCellNoWrap>
                <TableCell align="right" />
              </TableRow>
            </TableHead>
            <TableBody>
              {pipeds.map((piped) => (
                <PipedTableRow
                  key={piped.id}
                  pipedId={piped.id}
                  onEdit={handleEdit}
                  onDisable={handleDisable}
                  onEnable={handleEnable}
                  onRestart={handleRestart}
                />
              ))}
            </TableBody>
          </Table>
        </TableContainer>

        {openFilter && (
          <PipedFilter values={filterValues} onChange={setFilterValues} />
        )}
      </Box>
      <AddPipedDialog open={isOpenForm} onClose={handleClose} />
      <EditPipedDialog pipedId={editPipedId} onClose={handleEditClose} />
      <UpgradePipedDialog
        open={isUpgradeDialogOpen}
        pipeds={pipeds}
        releasedVersions={releasedVersions}
        onClose={handleUpgradeDialogClose}
      />
      <Dialog fullWidth open={Boolean(registeredPiped)}>
        <DialogTitle>
          {registeredPiped?.isNewKey
            ? "Added a new piped key"
            : "Piped registered"}
        </DialogTitle>
        {registeredPiped?.isNewKey ? (
          <Alert severity="info">{OLD_KEY_ALERT_MESSAGE}</Alert>
        ) : null}
        <DialogContent>
          <TextWithCopyButton
            name="Piped Id"
            value={registeredPiped?.id ?? ""}
          />
          <TextWithCopyButton
            name="Piped Key"
            value={registeredPiped?.key ?? ""}
          />
          <TextWithCopyButton
            name="Base64 Encoded Piped Key"
            value={
              registeredPiped?.key !== undefined
                ? btoa(registeredPiped?.key)
                : ""
            }
          />
          <Box
            sx={{
              display: "flex",
              justifyContent: "flex-end",
              m: 1,
              mt: 2,
            }}
          >
            <Button color="primary" onClick={handleClosePipedInfo}>
              {UI_TEXT_CLOSE}
            </Button>
          </Box>
        </DialogContent>
      </Dialog>
    </>
  );
});
