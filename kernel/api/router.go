// SiYuan - Refactor your thinking
// Copyright (c) 2020-present, b3log.org
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program.  If not, see <https://www.gnu.org/licenses/>.

package api

import (
	"github.com/gin-gonic/gin"
	"github.com/siyuan-community/siyuan/kernel/model"
)

func ServeAPI(ginServer *gin.Engine) {
	// 不需要鉴权

	ginServer.Handle("GET", "/api/system/bootProgress", bootProgress)
	ginServer.Handle("POST", "/api/system/bootProgress", bootProgress)
	ginServer.Handle("GET", "/api/system/version", version)
	ginServer.Handle("POST", "/api/system/version", version)
	ginServer.Handle("POST", "/api/system/currentTime", currentTime)
	ginServer.Handle("POST", "/api/system/uiproc", addUIProcess)
	ginServer.Handle("POST", "/api/system/loginAuth", model.LoginAuth)
	ginServer.Handle("POST", "/api/system/logoutAuth", model.LogoutAuth)
	ginServer.Handle("GET", "/api/system/getCaptcha", model.GetCaptcha)
	ginServer.Handle("POST", "/api/system/setUILayout", setUILayout) // 这里不加鉴权 After modifying the access authentication code on the browser side, the other side does not refresh https://github.com/siyuan-note/siyuan/issues/8028
	ginServer.Handle("GET", "/snippets/*filepath", serveSnippets)

	// 需要鉴权

	ginServer.Handle("POST", "/api/system/getEmojiConf", model.CheckAuth, getEmojiConf)
	ginServer.Handle("POST", "/api/system/setAPIToken", model.CheckAuth, model.CheckReadonly, setAPIToken)
	ginServer.Handle("POST", "/api/system/setAccessAuthCode", model.CheckAuth, model.CheckReadonly, setAccessAuthCode)
	ginServer.Handle("POST", "/api/system/setFollowSystemLockScreen", model.CheckAuth, model.CheckReadonly, setFollowSystemLockScreen)
	ginServer.Handle("POST", "/api/system/setNetworkServe", model.CheckAuth, model.CheckReadonly, setNetworkServe)
	ginServer.Handle("POST", "/api/system/setUploadErrLog", model.CheckAuth, model.CheckReadonly, setUploadErrLog)
	ginServer.Handle("POST", "/api/system/setAutoLaunch", model.CheckAuth, model.CheckReadonly, setAutoLaunch)
	ginServer.Handle("POST", "/api/system/setGoogleAnalytics", model.CheckAuth, model.CheckReadonly, setGoogleAnalytics)
	ginServer.Handle("POST", "/api/system/setDownloadInstallPkg", model.CheckAuth, model.CheckReadonly, setDownloadInstallPkg)
	ginServer.Handle("POST", "/api/system/setNetworkProxy", model.CheckAuth, model.CheckReadonly, setNetworkProxy)
	ginServer.Handle("POST", "/api/system/setWorkspaceDir", model.CheckAuth, model.CheckReadonly, setWorkspaceDir)
	ginServer.Handle("POST", "/api/system/getWorkspaces", model.CheckAuth, getWorkspaces)
	ginServer.Handle("POST", "/api/system/getMobileWorkspaces", model.CheckAuth, getMobileWorkspaces)
	ginServer.Handle("POST", "/api/system/checkWorkspaceDir", model.CheckAuth, model.CheckReadonly, checkWorkspaceDir)
	ginServer.Handle("POST", "/api/system/createWorkspaceDir", model.CheckAuth, model.CheckReadonly, createWorkspaceDir)
	ginServer.Handle("POST", "/api/system/removeWorkspaceDir", model.CheckAuth, model.CheckReadonly, removeWorkspaceDir)
	ginServer.Handle("POST", "/api/system/removeWorkspaceDirPhysically", model.CheckAuth, model.CheckReadonly, removeWorkspaceDirPhysically)
	ginServer.Handle("POST", "/api/system/setAppearanceMode", model.CheckAuth, setAppearanceMode)
	ginServer.Handle("POST", "/api/system/getSysFonts", model.CheckAuth, getSysFonts)
	ginServer.Handle("POST", "/api/system/exit", model.CheckAuth, exit)
	ginServer.Handle("POST", "/api/system/getConf", model.CheckAuth, getConf)
	ginServer.Handle("POST", "/api/system/checkUpdate", model.CheckAuth, checkUpdate)
	ginServer.Handle("POST", "/api/system/exportLog", model.CheckAuth, exportLog)
	ginServer.Handle("POST", "/api/system/getChangelog", model.CheckAuth, getChangelog)
	ginServer.Handle("POST", "/api/system/getNetwork", model.CheckAuth, getNetwork)

	ginServer.Handle("POST", "/api/storage/setLocalStorage", model.CheckAuth, model.CheckReadonly, setLocalStorage)
	ginServer.Handle("POST", "/api/storage/getLocalStorage", model.CheckAuth, getLocalStorage)
	ginServer.Handle("POST", "/api/storage/setLocalStorageVal", model.CheckAuth, model.CheckReadonly, setLocalStorageVal)
	ginServer.Handle("POST", "/api/storage/removeLocalStorageVals", model.CheckAuth, model.CheckReadonly, removeLocalStorageVals)
	ginServer.Handle("POST", "/api/storage/setCriterion", model.CheckAuth, model.CheckReadonly, setCriterion)
	ginServer.Handle("POST", "/api/storage/getCriteria", model.CheckAuth, getCriteria)
	ginServer.Handle("POST", "/api/storage/removeCriterion", model.CheckAuth, model.CheckReadonly, removeCriterion)
	ginServer.Handle("POST", "/api/storage/getRecentDocs", model.CheckAuth, getRecentDocs)

	ginServer.Handle("POST", "/api/account/login", model.CheckAuth, model.CheckReadonly, login)
	ginServer.Handle("POST", "/api/account/checkActivationcode", model.CheckAuth, model.CheckReadonly, checkActivationcode)
	ginServer.Handle("POST", "/api/account/useActivationcode", model.CheckAuth, model.CheckReadonly, useActivationcode)
	ginServer.Handle("POST", "/api/account/deactivate", model.CheckAuth, model.CheckReadonly, deactivateUser)
	ginServer.Handle("POST", "/api/account/startFreeTrial", model.CheckAuth, model.CheckReadonly, startFreeTrial)

	ginServer.Handle("POST", "/api/notebook/lsNotebooks", model.CheckAuth, lsNotebooks)
	ginServer.Handle("POST", "/api/notebook/openNotebook", model.CheckAuth, model.CheckReadonly, openNotebook)
	ginServer.Handle("POST", "/api/notebook/closeNotebook", model.CheckAuth, model.CheckReadonly, closeNotebook)
	ginServer.Handle("POST", "/api/notebook/getNotebookConf", model.CheckAuth, getNotebookConf)
	ginServer.Handle("POST", "/api/notebook/setNotebookConf", model.CheckAuth, model.CheckReadonly, setNotebookConf)
	ginServer.Handle("POST", "/api/notebook/createNotebook", model.CheckAuth, model.CheckReadonly, createNotebook)
	ginServer.Handle("POST", "/api/notebook/removeNotebook", model.CheckAuth, model.CheckReadonly, removeNotebook)
	ginServer.Handle("POST", "/api/notebook/renameNotebook", model.CheckAuth, model.CheckReadonly, renameNotebook)
	ginServer.Handle("POST", "/api/notebook/changeSortNotebook", model.CheckAuth, model.CheckReadonly, changeSortNotebook)
	ginServer.Handle("POST", "/api/notebook/setNotebookIcon", model.CheckAuth, model.CheckReadonly, setNotebookIcon)

	ginServer.Handle("POST", "/api/filetree/searchDocs", model.CheckAuth, searchDocs)
	ginServer.Handle("POST", "/api/filetree/listDocsByPath", model.CheckAuth, listDocsByPath)
	ginServer.Handle("POST", "/api/filetree/getDoc", model.CheckAuth, getDoc)
	ginServer.Handle("POST", "/api/filetree/getDocCreateSavePath", model.CheckAuth, getDocCreateSavePath)
	ginServer.Handle("POST", "/api/filetree/getRefCreateSavePath", model.CheckAuth, getRefCreateSavePath)
	ginServer.Handle("POST", "/api/filetree/changeSort", model.CheckAuth, model.CheckReadonly, changeSort)
	ginServer.Handle("POST", "/api/filetree/createDocWithMd", model.CheckAuth, model.CheckReadonly, createDocWithMd)
	ginServer.Handle("POST", "/api/filetree/createDailyNote", model.CheckAuth, model.CheckReadonly, createDailyNote)
	ginServer.Handle("POST", "/api/filetree/createDoc", model.CheckAuth, model.CheckReadonly, createDoc)
	ginServer.Handle("POST", "/api/filetree/renameDoc", model.CheckAuth, model.CheckReadonly, renameDoc)
	ginServer.Handle("POST", "/api/filetree/removeDoc", model.CheckAuth, model.CheckReadonly, removeDoc)
	ginServer.Handle("POST", "/api/filetree/removeDocs", model.CheckAuth, model.CheckReadonly, removeDocs)
	ginServer.Handle("POST", "/api/filetree/moveDocs", model.CheckAuth, model.CheckReadonly, moveDocs)
	ginServer.Handle("POST", "/api/filetree/duplicateDoc", model.CheckAuth, model.CheckReadonly, duplicateDoc)
	ginServer.Handle("POST", "/api/filetree/getHPathByPath", model.CheckAuth, getHPathByPath)
	ginServer.Handle("POST", "/api/filetree/getHPathsByPaths", model.CheckAuth, getHPathsByPaths)
	ginServer.Handle("POST", "/api/filetree/getHPathByID", model.CheckAuth, getHPathByID)
	ginServer.Handle("POST", "/api/filetree/getFullHPathByID", model.CheckAuth, getFullHPathByID)
	ginServer.Handle("POST", "/api/filetree/getIDsByHPath", model.CheckAuth, getIDsByHPath)
	ginServer.Handle("POST", "/api/filetree/doc2Heading", model.CheckAuth, model.CheckReadonly, doc2Heading)
	ginServer.Handle("POST", "/api/filetree/heading2Doc", model.CheckAuth, model.CheckReadonly, heading2Doc)
	ginServer.Handle("POST", "/api/filetree/li2Doc", model.CheckAuth, model.CheckReadonly, li2Doc)
	ginServer.Handle("POST", "/api/filetree/refreshFiletree", model.CheckAuth, model.CheckReadonly, refreshFiletree)
	ginServer.Handle("POST", "/api/filetree/upsertIndexes", model.CheckAuth, model.CheckReadonly, upsertIndexes)
	ginServer.Handle("POST", "/api/filetree/removeIndexes", model.CheckAuth, model.CheckReadonly, removeIndexes)
	ginServer.Handle("POST", "/api/filetree/listDocTree", model.CheckAuth, model.CheckReadonly, listDocTree)

	ginServer.Handle("POST", "/api/format/autoSpace", model.CheckAuth, model.CheckReadonly, autoSpace)
	ginServer.Handle("POST", "/api/format/netImg2LocalAssets", model.CheckAuth, model.CheckReadonly, netImg2LocalAssets)
	ginServer.Handle("POST", "/api/format/netAssets2LocalAssets", model.CheckAuth, model.CheckReadonly, netAssets2LocalAssets)

	ginServer.Handle("POST", "/api/history/getNotebookHistory", model.CheckAuth, getNotebookHistory)
	ginServer.Handle("POST", "/api/history/rollbackNotebookHistory", model.CheckAuth, model.CheckReadonly, rollbackNotebookHistory)
	ginServer.Handle("POST", "/api/history/rollbackAssetsHistory", model.CheckAuth, model.CheckReadonly, rollbackAssetsHistory)
	ginServer.Handle("POST", "/api/history/getDocHistoryContent", model.CheckAuth, getDocHistoryContent)
	ginServer.Handle("POST", "/api/history/rollbackDocHistory", model.CheckAuth, model.CheckReadonly, rollbackDocHistory)
	ginServer.Handle("POST", "/api/history/clearWorkspaceHistory", model.CheckAuth, model.CheckReadonly, clearWorkspaceHistory)
	ginServer.Handle("POST", "/api/history/reindexHistory", model.CheckAuth, model.CheckReadonly, reindexHistory)
	ginServer.Handle("POST", "/api/history/searchHistory", model.CheckAuth, searchHistory)
	ginServer.Handle("POST", "/api/history/getHistoryItems", model.CheckAuth, getHistoryItems)

	ginServer.Handle("POST", "/api/outline/getDocOutline", model.CheckAuth, getDocOutline)
	ginServer.Handle("POST", "/api/bookmark/getBookmark", model.CheckAuth, getBookmark)
	ginServer.Handle("POST", "/api/bookmark/renameBookmark", model.CheckAuth, model.CheckReadonly, renameBookmark)
	ginServer.Handle("POST", "/api/bookmark/removeBookmark", model.CheckAuth, model.CheckReadonly, removeBookmark)
	ginServer.Handle("POST", "/api/tag/getTag", model.CheckAuth, getTag)
	ginServer.Handle("POST", "/api/tag/renameTag", model.CheckAuth, model.CheckReadonly, renameTag)
	ginServer.Handle("POST", "/api/tag/removeTag", model.CheckAuth, model.CheckReadonly, removeTag)

	ginServer.Handle("POST", "/api/lute/spinBlockDOM", model.CheckAuth, spinBlockDOM) // 未测试
	ginServer.Handle("POST", "/api/lute/html2BlockDOM", model.CheckAuth, html2BlockDOM)
	ginServer.Handle("POST", "/api/lute/copyStdMarkdown", model.CheckAuth, copyStdMarkdown)

	ginServer.Handle("POST", "/api/query/sql", model.CheckAuth, SQL)
	ginServer.Handle("POST", "/api/sqlite/flushTransaction", model.CheckAuth, model.CheckReadonly, flushTransaction)

	ginServer.Handle("POST", "/api/search/searchTag", model.CheckAuth, searchTag)
	ginServer.Handle("POST", "/api/search/searchTemplate", model.CheckAuth, searchTemplate)
	ginServer.Handle("POST", "/api/search/removeTemplate", model.CheckAuth, model.CheckReadonly, removeTemplate)
	ginServer.Handle("POST", "/api/search/searchWidget", model.CheckAuth, searchWidget)
	ginServer.Handle("POST", "/api/search/searchRefBlock", model.CheckAuth, searchRefBlock)
	ginServer.Handle("POST", "/api/search/searchEmbedBlock", model.CheckAuth, searchEmbedBlock)
	ginServer.Handle("POST", "/api/search/getEmbedBlock", model.CheckAuth, getEmbedBlock)
	ginServer.Handle("POST", "/api/search/updateEmbedBlock", model.CheckAuth, updateEmbedBlock)
	ginServer.Handle("POST", "/api/search/fullTextSearchBlock", model.CheckAuth, fullTextSearchBlock)
	ginServer.Handle("POST", "/api/search/searchAsset", model.CheckAuth, searchAsset)
	ginServer.Handle("POST", "/api/search/findReplace", model.CheckAuth, findReplace)
	ginServer.Handle("POST", "/api/search/fullTextSearchAssetContent", model.CheckAuth, fullTextSearchAssetContent)
	ginServer.Handle("POST", "/api/search/getAssetContent", model.CheckAuth, getAssetContent)
	ginServer.Handle("POST", "/api/search/listInvalidBlockRefs", model.CheckAuth, listInvalidBlockRefs)

	ginServer.Handle("POST", "/api/block/getBlockInfo", model.CheckAuth, getBlockInfo)
	ginServer.Handle("POST", "/api/block/getBlockDOM", model.CheckAuth, getBlockDOM)
	ginServer.Handle("POST", "/api/block/getBlockKramdown", model.CheckAuth, getBlockKramdown)
	ginServer.Handle("POST", "/api/block/getChildBlocks", model.CheckAuth, getChildBlocks)
	ginServer.Handle("POST", "/api/block/getTailChildBlocks", model.CheckAuth, getTailChildBlocks)
	ginServer.Handle("POST", "/api/block/getBlockBreadcrumb", model.CheckAuth, getBlockBreadcrumb)
	ginServer.Handle("POST", "/api/block/getBlockIndex", model.CheckAuth, getBlockIndex)
	ginServer.Handle("POST", "/api/block/getBlocksIndexes", model.CheckAuth, getBlocksIndexes)
	ginServer.Handle("POST", "/api/block/getRefIDs", model.CheckAuth, getRefIDs)
	ginServer.Handle("POST", "/api/block/getRefIDsByFileAnnotationID", model.CheckAuth, getRefIDsByFileAnnotationID)
	ginServer.Handle("POST", "/api/block/getBlockDefIDsByRefText", model.CheckAuth, getBlockDefIDsByRefText)
	ginServer.Handle("POST", "/api/block/getRefText", model.CheckAuth, getRefText)
	ginServer.Handle("POST", "/api/block/getDOMText", model.CheckAuth, getDOMText)
	ginServer.Handle("POST", "/api/block/getTreeStat", model.CheckAuth, getTreeStat)
	ginServer.Handle("POST", "/api/block/getBlocksWordCount", model.CheckAuth, getBlocksWordCount)
	ginServer.Handle("POST", "/api/block/getContentWordCount", model.CheckAuth, getContentWordCount)
	ginServer.Handle("POST", "/api/block/getRecentUpdatedBlocks", model.CheckAuth, getRecentUpdatedBlocks)
	ginServer.Handle("POST", "/api/block/getDocInfo", model.CheckAuth, getDocInfo)
	ginServer.Handle("POST", "/api/block/checkBlockExist", model.CheckAuth, checkBlockExist)
	ginServer.Handle("POST", "/api/block/checkBlockFold", model.CheckAuth, checkBlockFold)
	ginServer.Handle("POST", "/api/block/insertBlock", model.CheckAuth, model.CheckReadonly, insertBlock)
	ginServer.Handle("POST", "/api/block/prependBlock", model.CheckAuth, model.CheckReadonly, prependBlock)
	ginServer.Handle("POST", "/api/block/appendBlock", model.CheckAuth, model.CheckReadonly, appendBlock)
	ginServer.Handle("POST", "/api/block/appendDailyNoteBlock", model.CheckAuth, model.CheckReadonly, appendDailyNoteBlock)
	ginServer.Handle("POST", "/api/block/prependDailyNoteBlock", model.CheckAuth, model.CheckReadonly, prependDailyNoteBlock)
	ginServer.Handle("POST", "/api/block/updateBlock", model.CheckAuth, model.CheckReadonly, updateBlock)
	ginServer.Handle("POST", "/api/block/deleteBlock", model.CheckAuth, model.CheckReadonly, deleteBlock)
	ginServer.Handle("POST", "/api/block/moveBlock", model.CheckAuth, model.CheckReadonly, moveBlock)
	ginServer.Handle("POST", "/api/block/moveOutlineHeading", model.CheckAuth, model.CheckReadonly, moveOutlineHeading)
	ginServer.Handle("POST", "/api/block/foldBlock", model.CheckAuth, model.CheckReadonly, foldBlock)
	ginServer.Handle("POST", "/api/block/unfoldBlock", model.CheckAuth, model.CheckReadonly, unfoldBlock)
	ginServer.Handle("POST", "/api/block/setBlockReminder", model.CheckAuth, model.CheckReadonly, setBlockReminder)
	ginServer.Handle("POST", "/api/block/getHeadingLevelTransaction", model.CheckAuth, getHeadingLevelTransaction)
	ginServer.Handle("POST", "/api/block/getHeadingDeleteTransaction", model.CheckAuth, getHeadingDeleteTransaction)
	ginServer.Handle("POST", "/api/block/getHeadingChildrenIDs", model.CheckAuth, getHeadingChildrenIDs)
	ginServer.Handle("POST", "/api/block/getHeadingChildrenDOM", model.CheckAuth, getHeadingChildrenDOM)
	ginServer.Handle("POST", "/api/block/swapBlockRef", model.CheckAuth, model.CheckReadonly, swapBlockRef)
	ginServer.Handle("POST", "/api/block/transferBlockRef", model.CheckAuth, model.CheckReadonly, transferBlockRef)
	ginServer.Handle("POST", "/api/block/getBlockSiblingID", model.CheckAuth, getBlockSiblingID)
	ginServer.Handle("POST", "/api/block/getBlockTreeInfos", model.CheckAuth, getBlockTreeInfos)

	ginServer.Handle("POST", "/api/file/getFile", model.CheckAuth, getFile)
	ginServer.Handle("POST", "/api/file/putFile", model.CheckAuth, model.CheckReadonly, putFile)
	ginServer.Handle("POST", "/api/file/copyFile", model.CheckAuth, model.CheckReadonly, copyFile)
	ginServer.Handle("POST", "/api/file/globalCopyFiles", model.CheckAuth, model.CheckReadonly, globalCopyFiles)
	ginServer.Handle("POST", "/api/file/removeFile", model.CheckAuth, model.CheckReadonly, removeFile)
	ginServer.Handle("POST", "/api/file/renameFile", model.CheckAuth, model.CheckReadonly, renameFile)
	ginServer.Handle("POST", "/api/file/readDir", model.CheckAuth, readDir)
	ginServer.Handle("POST", "/api/file/getUniqueFilename", model.CheckAuth, getUniqueFilename)

	ginServer.Handle("POST", "/api/ref/refreshBacklink", model.CheckAuth, refreshBacklink)
	ginServer.Handle("POST", "/api/ref/getBacklink", model.CheckAuth, getBacklink)
	ginServer.Handle("POST", "/api/ref/getBacklink2", model.CheckAuth, getBacklink2)
	ginServer.Handle("POST", "/api/ref/getBacklinkDoc", model.CheckAuth, getBacklinkDoc)
	ginServer.Handle("POST", "/api/ref/getBackmentionDoc", model.CheckAuth, getBackmentionDoc)

	ginServer.Handle("POST", "/api/attr/getBookmarkLabels", model.CheckAuth, getBookmarkLabels)
	ginServer.Handle("POST", "/api/attr/resetBlockAttrs", model.CheckAuth, model.CheckReadonly, resetBlockAttrs)
	ginServer.Handle("POST", "/api/attr/setBlockAttrs", model.CheckAuth, model.CheckReadonly, setBlockAttrs)
	ginServer.Handle("POST", "/api/attr/batchSetBlockAttrs", model.CheckAuth, model.CheckReadonly, batchSetBlockAttrs)
	ginServer.Handle("POST", "/api/attr/getBlockAttrs", model.CheckAuth, getBlockAttrs)

	ginServer.Handle("POST", "/api/cloud/getCloudSpace", model.CheckAuth, getCloudSpace)

	ginServer.Handle("POST", "/api/sync/setSyncEnable", model.CheckAuth, model.CheckReadonly, setSyncEnable)
	ginServer.Handle("POST", "/api/sync/setSyncPerception", model.CheckAuth, model.CheckReadonly, setSyncPerception)
	ginServer.Handle("POST", "/api/sync/setSyncGenerateConflictDoc", model.CheckAuth, model.CheckReadonly, setSyncGenerateConflictDoc)
	ginServer.Handle("POST", "/api/sync/setSyncMode", model.CheckAuth, model.CheckReadonly, setSyncMode)
	ginServer.Handle("POST", "/api/sync/setSyncProvider", model.CheckAuth, model.CheckReadonly, setSyncProvider)
	ginServer.Handle("POST", "/api/sync/setSyncProviderS3", model.CheckAuth, model.CheckReadonly, setSyncProviderS3)
	ginServer.Handle("POST", "/api/sync/setSyncProviderWebDAV", model.CheckAuth, model.CheckReadonly, setSyncProviderWebDAV)
	ginServer.Handle("POST", "/api/sync/setCloudSyncDir", model.CheckAuth, model.CheckReadonly, setCloudSyncDir)
	ginServer.Handle("POST", "/api/sync/createCloudSyncDir", model.CheckAuth, model.CheckReadonly, createCloudSyncDir)
	ginServer.Handle("POST", "/api/sync/removeCloudSyncDir", model.CheckAuth, model.CheckReadonly, removeCloudSyncDir)
	ginServer.Handle("POST", "/api/sync/listCloudSyncDir", model.CheckAuth, listCloudSyncDir)
	ginServer.Handle("POST", "/api/sync/performSync", model.CheckAuth, model.CheckReadonly, performSync)
	ginServer.Handle("POST", "/api/sync/performBootSync", model.CheckAuth, model.CheckReadonly, performBootSync)
	ginServer.Handle("POST", "/api/sync/getBootSync", model.CheckAuth, getBootSync)
	ginServer.Handle("POST", "/api/sync/getSyncInfo", model.CheckAuth, getSyncInfo)
	ginServer.Handle("POST", "/api/sync/exportSyncProviderS3", model.CheckAuth, exportSyncProviderS3)
	ginServer.Handle("POST", "/api/sync/importSyncProviderS3", model.CheckAuth, model.CheckReadonly, importSyncProviderS3)
	ginServer.Handle("POST", "/api/sync/exportSyncProviderWebDAV", model.CheckAuth, exportSyncProviderWebDAV)
	ginServer.Handle("POST", "/api/sync/importSyncProviderWebDAV", model.CheckAuth, model.CheckReadonly, importSyncProviderWebDAV)

	ginServer.Handle("POST", "/api/inbox/getShorthands", model.CheckAuth, getShorthands)
	ginServer.Handle("POST", "/api/inbox/getShorthand", model.CheckAuth, getShorthand)
	ginServer.Handle("POST", "/api/inbox/removeShorthands", model.CheckAuth, model.CheckReadonly, removeShorthands)

	ginServer.Handle("POST", "/api/extension/copy", model.CheckAuth, model.CheckReadonly, extensionCopy)

	ginServer.Handle("POST", "/api/clipboard/readFilePaths", model.CheckAuth, readFilePaths)

	ginServer.Handle("POST", "/api/asset/uploadCloud", model.CheckAuth, model.CheckReadonly, uploadCloud)
	ginServer.Handle("POST", "/api/asset/insertLocalAssets", model.CheckAuth, model.CheckReadonly, insertLocalAssets)
	ginServer.Handle("POST", "/api/asset/resolveAssetPath", model.CheckAuth, resolveAssetPath)
	ginServer.Handle("POST", "/api/asset/upload", model.CheckAuth, model.CheckReadonly, model.Upload)
	ginServer.Handle("POST", "/api/asset/setFileAnnotation", model.CheckAuth, model.CheckReadonly, setFileAnnotation)
	ginServer.Handle("POST", "/api/asset/getFileAnnotation", model.CheckAuth, getFileAnnotation)
	ginServer.Handle("POST", "/api/asset/getUnusedAssets", model.CheckAuth, getUnusedAssets)
	ginServer.Handle("POST", "/api/asset/getMissingAssets", model.CheckAuth, getMissingAssets)
	ginServer.Handle("POST", "/api/asset/removeUnusedAsset", model.CheckAuth, model.CheckReadonly, removeUnusedAsset)
	ginServer.Handle("POST", "/api/asset/removeUnusedAssets", model.CheckAuth, model.CheckReadonly, removeUnusedAssets)
	ginServer.Handle("POST", "/api/asset/getDocImageAssets", model.CheckAuth, getDocImageAssets)
	ginServer.Handle("POST", "/api/asset/renameAsset", model.CheckAuth, model.CheckReadonly, renameAsset)
	ginServer.Handle("POST", "/api/asset/getImageOCRText", model.CheckAuth, model.CheckReadonly, getImageOCRText)
	ginServer.Handle("POST", "/api/asset/setImageOCRText", model.CheckAuth, model.CheckReadonly, setImageOCRText)
	ginServer.Handle("POST", "/api/asset/fullReindexAssetContent", model.CheckAuth, model.CheckReadonly, fullReindexAssetContent)
	ginServer.Handle("POST", "/api/asset/statAsset", model.CheckAuth, statAsset)

	ginServer.Handle("POST", "/api/export/batchExportMd", model.CheckAuth, batchExportMd)
	ginServer.Handle("POST", "/api/export/exportMd", model.CheckAuth, exportMd)
	ginServer.Handle("POST", "/api/export/exportSY", model.CheckAuth, exportSY)
	ginServer.Handle("POST", "/api/export/exportNotebookSY", model.CheckAuth, exportNotebookSY)
	ginServer.Handle("POST", "/api/export/exportMdContent", model.CheckAuth, exportMdContent)
	ginServer.Handle("POST", "/api/export/exportHTML", model.CheckAuth, exportHTML)
	ginServer.Handle("POST", "/api/export/exportPreviewHTML", model.CheckAuth, exportPreviewHTML)
	ginServer.Handle("POST", "/api/export/exportMdHTML", model.CheckAuth, exportMdHTML)
	ginServer.Handle("POST", "/api/export/exportDocx", model.CheckAuth, exportDocx)
	ginServer.Handle("POST", "/api/export/processPDF", model.CheckAuth, processPDF)
	ginServer.Handle("POST", "/api/export/preview", model.CheckAuth, exportPreview)
	ginServer.Handle("POST", "/api/export/exportResources", model.CheckAuth, exportResources)
	ginServer.Handle("POST", "/api/export/exportAsFile", model.CheckAuth, exportAsFile)
	ginServer.Handle("POST", "/api/export/exportData", model.CheckAuth, exportData)
	ginServer.Handle("POST", "/api/export/exportDataInFolder", model.CheckAuth, exportDataInFolder)
	ginServer.Handle("POST", "/api/export/exportTempContent", model.CheckAuth, exportTempContent)
	ginServer.Handle("POST", "/api/export/export2Liandi", model.CheckAuth, export2Liandi)
	ginServer.Handle("POST", "/api/export/exportReStructuredText", model.CheckAuth, exportReStructuredText)
	ginServer.Handle("POST", "/api/export/exportAsciiDoc", model.CheckAuth, exportAsciiDoc)
	ginServer.Handle("POST", "/api/export/exportTextile", model.CheckAuth, exportTextile)
	ginServer.Handle("POST", "/api/export/exportOPML", model.CheckAuth, exportOPML)
	ginServer.Handle("POST", "/api/export/exportOrgMode", model.CheckAuth, exportOrgMode)
	ginServer.Handle("POST", "/api/export/exportMediaWiki", model.CheckAuth, exportMediaWiki)
	ginServer.Handle("POST", "/api/export/exportODT", model.CheckAuth, exportODT)
	ginServer.Handle("POST", "/api/export/exportRTF", model.CheckAuth, exportRTF)
	ginServer.Handle("POST", "/api/export/exportEPUB", model.CheckAuth, exportEPUB)
	ginServer.Handle("POST", "/api/export/exportAttributeView", model.CheckAuth, exportAttributeView)

	ginServer.Handle("POST", "/api/import/importStdMd", model.CheckAuth, model.CheckReadonly, importStdMd)
	ginServer.Handle("POST", "/api/import/importData", model.CheckAuth, model.CheckReadonly, importData)
	ginServer.Handle("POST", "/api/import/importSY", model.CheckAuth, model.CheckReadonly, importSY)

	ginServer.Handle("POST", "/api/convert/pandoc", model.CheckAuth, model.CheckReadonly, pandoc)

	ginServer.Handle("POST", "/api/template/render", model.CheckAuth, renderTemplate)
	ginServer.Handle("POST", "/api/template/docSaveAsTemplate", model.CheckAuth, model.CheckReadonly, docSaveAsTemplate)
	ginServer.Handle("POST", "/api/template/renderSprig", model.CheckAuth, renderSprig)

	ginServer.Handle("POST", "/api/transactions", model.CheckAuth, model.CheckReadonly, performTransactions)

	ginServer.Handle("POST", "/api/setting/setAccount", model.CheckAuth, model.CheckReadonly, setAccount)
	ginServer.Handle("POST", "/api/setting/setEditor", model.CheckAuth, model.CheckReadonly, setEditor)
	ginServer.Handle("POST", "/api/setting/setExport", model.CheckAuth, model.CheckReadonly, setExport)
	ginServer.Handle("POST", "/api/setting/setFiletree", model.CheckAuth, model.CheckReadonly, setFiletree)
	ginServer.Handle("POST", "/api/setting/setSearch", model.CheckAuth, model.CheckReadonly, setSearch)
	ginServer.Handle("POST", "/api/setting/setKeymap", model.CheckAuth, model.CheckReadonly, setKeymap)
	ginServer.Handle("POST", "/api/setting/setAppearance", model.CheckAuth, model.CheckReadonly, setAppearance)
	ginServer.Handle("POST", "/api/setting/getCloudUser", model.CheckAuth, getCloudUser)
	ginServer.Handle("POST", "/api/setting/logoutCloudUser", model.CheckAuth, model.CheckReadonly, logoutCloudUser)
	ginServer.Handle("POST", "/api/setting/login2faCloudUser", model.CheckAuth, model.CheckReadonly, login2faCloudUser)
	ginServer.Handle("POST", "/api/setting/setEmoji", model.CheckAuth, model.CheckReadonly, setEmoji)
	ginServer.Handle("POST", "/api/setting/setFlashcard", model.CheckAuth, model.CheckReadonly, setFlashcard)
	ginServer.Handle("POST", "/api/setting/setAI", model.CheckAuth, model.CheckReadonly, setAI)
	ginServer.Handle("POST", "/api/setting/setBazaar", model.CheckAuth, model.CheckReadonly, setBazaar)
	ginServer.Handle("POST", "/api/setting/refreshVirtualBlockRef", model.CheckAuth, model.CheckReadonly, refreshVirtualBlockRef)
	ginServer.Handle("POST", "/api/setting/addVirtualBlockRefInclude", model.CheckAuth, model.CheckReadonly, addVirtualBlockRefInclude)
	ginServer.Handle("POST", "/api/setting/addVirtualBlockRefExclude", model.CheckAuth, model.CheckReadonly, addVirtualBlockRefExclude)
	ginServer.Handle("POST", "/api/setting/setSnippet", model.CheckAuth, model.CheckReadonly, setConfSnippet)
	ginServer.Handle("POST", "/api/setting/setEditorReadOnly", model.CheckAuth, model.CheckReadonly, setEditorReadOnly)

	ginServer.Handle("POST", "/api/graph/resetGraph", model.CheckAuth, model.CheckReadonly, resetGraph)
	ginServer.Handle("POST", "/api/graph/resetLocalGraph", model.CheckAuth, model.CheckReadonly, resetLocalGraph)
	ginServer.Handle("POST", "/api/graph/getGraph", model.CheckAuth, getGraph)
	ginServer.Handle("POST", "/api/graph/getLocalGraph", model.CheckAuth, getLocalGraph)

	ginServer.Handle("POST", "/api/bazaar/getBazaarPlugin", model.CheckAuth, getBazaarPlugin)
	ginServer.Handle("POST", "/api/bazaar/getInstalledPlugin", model.CheckAuth, getInstalledPlugin)
	ginServer.Handle("POST", "/api/bazaar/installBazaarPlugin", model.CheckAuth, model.CheckReadonly, installBazaarPlugin)
	ginServer.Handle("POST", "/api/bazaar/uninstallBazaarPlugin", model.CheckAuth, model.CheckReadonly, uninstallBazaarPlugin)
	ginServer.Handle("POST", "/api/bazaar/getBazaarWidget", model.CheckAuth, getBazaarWidget)
	ginServer.Handle("POST", "/api/bazaar/getInstalledWidget", model.CheckAuth, getInstalledWidget)
	ginServer.Handle("POST", "/api/bazaar/installBazaarWidget", model.CheckAuth, model.CheckReadonly, installBazaarWidget)
	ginServer.Handle("POST", "/api/bazaar/uninstallBazaarWidget", model.CheckAuth, model.CheckReadonly, uninstallBazaarWidget)
	ginServer.Handle("POST", "/api/bazaar/getBazaarIcon", model.CheckAuth, getBazaarIcon)
	ginServer.Handle("POST", "/api/bazaar/getInstalledIcon", model.CheckAuth, getInstalledIcon)
	ginServer.Handle("POST", "/api/bazaar/installBazaarIcon", model.CheckAuth, model.CheckReadonly, installBazaarIcon)
	ginServer.Handle("POST", "/api/bazaar/uninstallBazaarIcon", model.CheckAuth, model.CheckReadonly, uninstallBazaarIcon)
	ginServer.Handle("POST", "/api/bazaar/getBazaarTemplate", model.CheckAuth, getBazaarTemplate)
	ginServer.Handle("POST", "/api/bazaar/getInstalledTemplate", model.CheckAuth, getInstalledTemplate)
	ginServer.Handle("POST", "/api/bazaar/installBazaarTemplate", model.CheckAuth, model.CheckReadonly, installBazaarTemplate)
	ginServer.Handle("POST", "/api/bazaar/uninstallBazaarTemplate", model.CheckAuth, model.CheckReadonly, uninstallBazaarTemplate)
	ginServer.Handle("POST", "/api/bazaar/getBazaarTheme", model.CheckAuth, getBazaarTheme)
	ginServer.Handle("POST", "/api/bazaar/getInstalledTheme", model.CheckAuth, getInstalledTheme)
	ginServer.Handle("POST", "/api/bazaar/installBazaarTheme", model.CheckAuth, model.CheckReadonly, installBazaarTheme)
	ginServer.Handle("POST", "/api/bazaar/uninstallBazaarTheme", model.CheckAuth, model.CheckReadonly, uninstallBazaarTheme)
	ginServer.Handle("POST", "/api/bazaar/getBazaarPackageREAME", model.CheckAuth, getBazaarPackageREAME)
	ginServer.Handle("POST", "/api/bazaar/getUpdatedPackage", model.CheckAuth, getUpdatedPackage)
	ginServer.Handle("POST", "/api/bazaar/batchUpdatePackage", model.CheckAuth, batchUpdatePackage)

	ginServer.Handle("POST", "/api/repo/initRepoKey", model.CheckAuth, model.CheckReadonly, initRepoKey)
	ginServer.Handle("POST", "/api/repo/initRepoKeyFromPassphrase", model.CheckAuth, model.CheckReadonly, initRepoKeyFromPassphrase)
	ginServer.Handle("POST", "/api/repo/resetRepo", model.CheckAuth, model.CheckReadonly, resetRepo)
	ginServer.Handle("POST", "/api/repo/purgeRepo", model.CheckAuth, model.CheckReadonly, purgeRepo)
	ginServer.Handle("POST", "/api/repo/purgeCloudRepo", model.CheckAuth, model.CheckReadonly, purgeCloudRepo)
	ginServer.Handle("POST", "/api/repo/importRepoKey", model.CheckAuth, model.CheckReadonly, importRepoKey)
	ginServer.Handle("POST", "/api/repo/createSnapshot", model.CheckAuth, model.CheckReadonly, createSnapshot)
	ginServer.Handle("POST", "/api/repo/tagSnapshot", model.CheckAuth, model.CheckReadonly, tagSnapshot)
	ginServer.Handle("POST", "/api/repo/checkoutRepo", model.CheckAuth, model.CheckReadonly, checkoutRepo)
	ginServer.Handle("POST", "/api/repo/getRepoSnapshots", model.CheckAuth, getRepoSnapshots)
	ginServer.Handle("POST", "/api/repo/getRepoTagSnapshots", model.CheckAuth, getRepoTagSnapshots)
	ginServer.Handle("POST", "/api/repo/removeRepoTagSnapshot", model.CheckAuth, model.CheckReadonly, removeRepoTagSnapshot)
	ginServer.Handle("POST", "/api/repo/getCloudRepoTagSnapshots", model.CheckAuth, getCloudRepoTagSnapshots)
	ginServer.Handle("POST", "/api/repo/getCloudRepoSnapshots", model.CheckAuth, getCloudRepoSnapshots)
	ginServer.Handle("POST", "/api/repo/removeCloudRepoTagSnapshot", model.CheckAuth, model.CheckReadonly, removeCloudRepoTagSnapshot)
	ginServer.Handle("POST", "/api/repo/uploadCloudSnapshot", model.CheckAuth, model.CheckReadonly, uploadCloudSnapshot)
	ginServer.Handle("POST", "/api/repo/downloadCloudSnapshot", model.CheckAuth, model.CheckReadonly, downloadCloudSnapshot)
	ginServer.Handle("POST", "/api/repo/diffRepoSnapshots", model.CheckAuth, diffRepoSnapshots)
	ginServer.Handle("POST", "/api/repo/openRepoSnapshotDoc", model.CheckAuth, openRepoSnapshotDoc)
	ginServer.Handle("POST", "/api/repo/getRepoFile", model.CheckAuth, getRepoFile)

	ginServer.Handle("POST", "/api/riff/createRiffDeck", model.CheckAuth, model.CheckReadonly, createRiffDeck)
	ginServer.Handle("POST", "/api/riff/renameRiffDeck", model.CheckAuth, model.CheckReadonly, renameRiffDeck)
	ginServer.Handle("POST", "/api/riff/removeRiffDeck", model.CheckAuth, model.CheckReadonly, removeRiffDeck)
	ginServer.Handle("POST", "/api/riff/getRiffDecks", model.CheckAuth, getRiffDecks)
	ginServer.Handle("POST", "/api/riff/addRiffCards", model.CheckAuth, model.CheckReadonly, addRiffCards)
	ginServer.Handle("POST", "/api/riff/removeRiffCards", model.CheckAuth, model.CheckReadonly, removeRiffCards)
	ginServer.Handle("POST", "/api/riff/getRiffDueCards", model.CheckAuth, getRiffDueCards)
	ginServer.Handle("POST", "/api/riff/getTreeRiffDueCards", model.CheckAuth, getTreeRiffDueCards)
	ginServer.Handle("POST", "/api/riff/getNotebookRiffDueCards", model.CheckAuth, getNotebookRiffDueCards)
	ginServer.Handle("POST", "/api/riff/reviewRiffCard", model.CheckAuth, model.CheckReadonly, reviewRiffCard)
	ginServer.Handle("POST", "/api/riff/skipReviewRiffCard", model.CheckAuth, model.CheckReadonly, skipReviewRiffCard)
	ginServer.Handle("POST", "/api/riff/getRiffCards", model.CheckAuth, getRiffCards)
	ginServer.Handle("POST", "/api/riff/getTreeRiffCards", model.CheckAuth, getTreeRiffCards)
	ginServer.Handle("POST", "/api/riff/getNotebookRiffCards", model.CheckAuth, getNotebookRiffCards)
	ginServer.Handle("POST", "/api/riff/resetRiffCards", model.CheckAuth, model.CheckReadonly, resetRiffCards)
	ginServer.Handle("POST", "/api/riff/batchSetRiffCardsDueTime", model.CheckAuth, model.CheckReadonly, batchSetRiffCardsDueTime)
	ginServer.Handle("POST", "/api/riff/getRiffCardsByBlockIDs", model.CheckAuth, model.CheckReadonly, getRiffCardsByBlockIDs)

	ginServer.Handle("POST", "/api/notification/pushMsg", model.CheckAuth, pushMsg)
	ginServer.Handle("POST", "/api/notification/pushErrMsg", model.CheckAuth, pushErrMsg)

	ginServer.Handle("POST", "/api/snippet/getSnippet", model.CheckAuth, getSnippet)
	ginServer.Handle("POST", "/api/snippet/setSnippet", model.CheckAuth, model.CheckReadonly, setSnippet)
	ginServer.Handle("POST", "/api/snippet/removeSnippet", model.CheckAuth, model.CheckReadonly, removeSnippet)

	ginServer.Handle("POST", "/api/av/renderAttributeView", model.CheckAuth, renderAttributeView)
	ginServer.Handle("POST", "/api/av/renderHistoryAttributeView", model.CheckAuth, renderHistoryAttributeView)
	ginServer.Handle("POST", "/api/av/renderSnapshotAttributeView", model.CheckAuth, renderSnapshotAttributeView)
	ginServer.Handle("POST", "/api/av/getAttributeViewKeys", model.CheckAuth, getAttributeViewKeys)
	ginServer.Handle("POST", "/api/av/setAttributeViewBlockAttr", model.CheckAuth, model.CheckReadonly, setAttributeViewBlockAttr)
	ginServer.Handle("POST", "/api/av/searchAttributeView", model.CheckAuth, model.CheckReadonly, searchAttributeView)
	ginServer.Handle("POST", "/api/av/getAttributeView", model.CheckAuth, model.CheckReadonly, getAttributeView)
	ginServer.Handle("POST", "/api/av/searchAttributeViewRelationKey", model.CheckAuth, model.CheckReadonly, searchAttributeViewRelationKey)
	ginServer.Handle("POST", "/api/av/searchAttributeViewNonRelationKey", model.CheckAuth, model.CheckReadonly, searchAttributeViewNonRelationKey)
	ginServer.Handle("POST", "/api/av/getAttributeViewFilterSort", model.CheckAuth, model.CheckReadonly, getAttributeViewFilterSort)
	ginServer.Handle("POST", "/api/av/addAttributeViewKey", model.CheckAuth, model.CheckReadonly, addAttributeViewKey)
	ginServer.Handle("POST", "/api/av/removeAttributeViewKey", model.CheckAuth, model.CheckReadonly, removeAttributeViewKey)
	ginServer.Handle("POST", "/api/av/sortAttributeViewViewKey", model.CheckAuth, model.CheckReadonly, sortAttributeViewViewKey)
	ginServer.Handle("POST", "/api/av/sortAttributeViewKey", model.CheckAuth, model.CheckReadonly, sortAttributeViewKey)
	ginServer.Handle("POST", "/api/av/addAttributeViewValues", model.CheckAuth, model.CheckReadonly, addAttributeViewValues)
	ginServer.Handle("POST", "/api/av/removeAttributeViewValues", model.CheckAuth, model.CheckReadonly, removeAttributeViewValues)
	ginServer.Handle("POST", "/api/av/getAttributeViewPrimaryKeyValues", model.CheckAuth, model.CheckReadonly, getAttributeViewPrimaryKeyValues)
	ginServer.Handle("POST", "/api/av/setDatabaseBlockView", model.CheckAuth, model.CheckReadonly, setDatabaseBlockView)
	ginServer.Handle("POST", "/api/av/getMirrorDatabaseBlocks", model.CheckAuth, model.CheckReadonly, getMirrorDatabaseBlocks)
	ginServer.Handle("POST", "/api/av/getAttributeViewKeysByAvID", model.CheckAuth, model.CheckReadonly, getAttributeViewKeysByAvID)

	ginServer.Handle("POST", "/api/ai/chatGPT", model.CheckAuth, chatGPT)
	ginServer.Handle("POST", "/api/ai/chatGPTWithAction", model.CheckAuth, chatGPTWithAction)

	ginServer.Handle("POST", "/api/petal/loadPetals", model.CheckAuth, loadPetals)
	ginServer.Handle("POST", "/api/petal/setPetalEnabled", model.CheckAuth, model.CheckReadonly, setPetalEnabled)

	ginServer.Any("/api/network/echo", model.CheckAuth, echo)
	ginServer.Handle("POST", "/api/network/forwardProxy", model.CheckAuth, forwardProxy)

	ginServer.Handle("GET", "/ws/broadcast", model.CheckAuth, broadcast)
	ginServer.Handle("POST", "/api/broadcast/postMessage", model.CheckAuth, postMessage)
	ginServer.Handle("POST", "/api/broadcast/getChannels", model.CheckAuth, getChannels)
	ginServer.Handle("POST", "/api/broadcast/getChannelInfo", model.CheckAuth, getChannelInfo)

	ginServer.Handle("POST", "/api/archive/zip", model.CheckAuth, model.CheckReadonly, zip)
	ginServer.Handle("POST", "/api/archive/unzip", model.CheckAuth, model.CheckReadonly, unzip)
}
